package auth

import (
	"context"
	"errors"
	"github.com/Siroshun09/logs"
	"github.com/Siroshun09/serrors"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/okocraft/monitor/internal/domain/auditlog"
	"github.com/okocraft/monitor/internal/domain/auth"
	"github.com/okocraft/monitor/internal/domain/user"
	"github.com/okocraft/monitor/internal/handler/oapi"
	auth2 "github.com/okocraft/monitor/internal/usecases/auth"
	"github.com/okocraft/monitor/internal/usecases/permission"
	user2 "github.com/okocraft/monitor/internal/usecases/user"
	"github.com/okocraft/monitor/lib/ctxlib"
	"github.com/okocraft/monitor/lib/httplib"
	"net/http"
	"strings"
	"time"
)

type AuthHandler struct {
	authUsecase       auth2.AuthUsecase
	userUsecase       user2.UserUsecase
	permissionUsecase permission.PermissionUsecase
}

func NewAuthHandler(authUsecase auth2.AuthUsecase, userUsecase user2.UserUsecase, permissionUseCase permission.PermissionUsecase) AuthHandler {
	return AuthHandler{
		authUsecase:       authUsecase,
		userUsecase:       userUsecase,
		permissionUsecase: permissionUseCase,
	}
}

func (h AuthHandler) SetAuthMethodIntoContext(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	var method auth.Method
	switch input.SecuritySchemeName {
	case "SkipAuth":
		method = auth.MethodSkip
	case "AccessTokenAuth":
		method = auth.MethodAccessToken
	default:
		logs.Warn(ctx, serrors.New("unknown auth method: "+input.SecuritySchemeName))
	}

	ctxlib.GetHTTPAccessLog(ctx).AuthMethod = method

	return nil
}

func (h AuthHandler) NewAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		accessLog := ctxlib.GetHTTPAccessLog(ctx)

		switch accessLog.AuthMethod {
		case auth.MethodSkip:
			next.ServeHTTP(w, r)
		case auth.MethodAccessToken:
			if userID, ok := h.AuthorizeByAccessToken(ctx, w, r); ok {
				accessLog.UserID = userID
				r = r.WithContext(ctxlib.WithUserID(ctx, userID))
				ctxlib.SetUserIDForAuditLog(ctx, userID)
				next.ServeHTTP(w, r)
			}
		default:
			httplib.RenderError(ctx, w, serrors.New("unknown auth method"))
		}
	})
}

func (h AuthHandler) AuthorizeByAccessToken(ctx context.Context, w http.ResponseWriter, r *http.Request) (user.ID, bool) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		httplib.RenderUnauthorized(ctx, w, serrors.New("no authorization header"))
		return 0, false
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	userID, err := h.authUsecase.VerifyAccessToken(ctx, tokenStr)

	switch {
	case auth.IsUnauthorizedError(err):
		httplib.RenderUnauthorized(ctx, w, err)
		return 0, false
	case err != nil:
		httplib.RenderError(ctx, w, err)
		return 0, false
	}

	return userID, true
}

func (h AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		httplib.RenderNoContent(ctx, w)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Path:     "/auth",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   -1,
	})

	userID, refreshTokenID, _, err := h.authUsecase.VerifyRefreshToken(ctx, cookie.Value)
	if auth.IsUnauthorizedError(err) {
		httplib.RenderUnauthorized(ctx, w, err)
		return
	} else if err != nil {
		httplib.RenderError(ctx, w, err)
		return
	}

	err = h.authUsecase.InvalidateTokens(ctx, refreshTokenID)
	if err != nil {
		httplib.RenderError(ctx, w, err)
		return
	}

	httplib.RenderNoContent(ctx, w)

	ctxlib.SetUserIDForAuditLog(ctx, userID)
	ctxlib.AddAuditLogRecord(ctx, auditlog.UserLogRecord{
		Action:    auditlog.UserActionLogout,
		Timestamp: time.Now(),
	})
}

func (h AuthHandler) RefreshAccessToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		httplib.RenderUnauthorized(ctx, w, err)
		return
	}

	userID, refreshTokenID, expiresAt, err := h.authUsecase.VerifyRefreshToken(ctx, cookie.Value)
	if auth.IsUnauthorizedError(err) {
		httplib.RenderUnauthorized(ctx, w, err)
		return
	} else if err != nil {
		httplib.RenderError(ctx, w, err)
		return
	}

	accessToken, err := h.authUsecase.RefreshAccessToken(ctx, userID, refreshTokenID, expiresAt)
	if err != nil {
		httplib.RenderError(ctx, w, err)
		return
	}

	ctx = ctxlib.WithUserID(ctx, userID)
	me, err := h.userUsecase.GetMe(ctx)
	var notFound user.NotFoundByIDError
	if errors.As(err, &notFound) {
		httplib.RenderNotFound(ctx, w, err)
		return
	} else if err != nil {
		httplib.RenderError(ctx, w, err)
		return
	}

	pagePermissions, err := h.permissionUsecase.CalculatePagePermissions(ctx)
	if err != nil {
		httplib.RenderError(ctx, w, err)
		return
	}

	httplib.RenderOK(ctx, w, oapi.AccessTokenWithMeAndPagePermissions{
		AccessToken:     accessToken,
		Me:              me.ToResponse(),
		PagePermissions: pagePermissions.ToResponse(),
	})
}
