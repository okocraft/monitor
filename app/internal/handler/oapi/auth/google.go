package auth

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Siroshun09/logs"
	"github.com/okocraft/monitor/internal/config"
	"github.com/okocraft/monitor/internal/domain/auditlog"
	"github.com/okocraft/monitor/internal/domain/auth"
	"github.com/okocraft/monitor/internal/domain/user"
	"github.com/okocraft/monitor/internal/handler/oapi"
	"github.com/okocraft/monitor/internal/usecases"
	"github.com/okocraft/monitor/lib/ctxlib"
	"github.com/okocraft/monitor/lib/httplib"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

type GoogleAuthHandler struct {
	enabled       bool
	resultPageURL string
	conf          oauth2.Config
	authUsecase   usecases.AuthUsecase
	userUsecase   usecases.UserUsecase
}

func NewGoogleAuthHandler(c config.GoogleAuthConfig, authUsecase usecases.AuthUsecase, userUsecase usecases.UserUsecase) GoogleAuthHandler {
	return GoogleAuthHandler{
		enabled:       c.Enabled,
		resultPageURL: c.ResultPageURL,
		conf: oauth2.Config{
			ClientID:     c.ClientID,
			ClientSecret: c.ClientSecret,
			RedirectURL:  c.RedirectURL,
			Scopes:       []string{"openid"},
			Endpoint:     google.Endpoint,
		},
		authUsecase: authUsecase,
		userUsecase: userUsecase,
	}
}

var (
	pkceKeyMap sync.Map
)

func (h GoogleAuthHandler) LinkWithGoogle(w http.ResponseWriter, r *http.Request, loginKey string) {
	ctx := r.Context()

	if !h.enabled {
		httplib.RenderNotAcceptable(ctx, w, nil)
		return
	}

	id, stateJWT, err := h.authUsecase.CreateStateJWTWithLoginKey(ctx, loginKey)
	if err != nil {
		httplib.RenderError(ctx, w, err)
	}

	verifier := oauth2.GenerateVerifier()
	pkceKeyMap.Store(id, verifier)

	redirectURL := h.conf.AuthCodeURL(stateJWT, oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))
	httplib.RenderOK(ctx, w, oapi.GoogleLoginURL{RedirectUrl: redirectURL})
}

func (h GoogleAuthHandler) LoginWithGoogle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if !h.enabled {
		httplib.RenderNotAcceptable(ctx, w, nil)
		return
	}

	currentPage, err := httplib.DecodeBody[oapi.CurrentPage](r)
	if err != nil {
		httplib.RenderBadRequest(ctx, w, err)
		return
	}

	id, stateJWT, err := h.authUsecase.CreateStateJWT(ctx, currentPage.Url)
	if err != nil {
		httplib.RenderError(ctx, w, err)
		return
	}

	verifier := oauth2.GenerateVerifier()
	pkceKeyMap.Store(id, verifier)

	redirectURL := h.conf.AuthCodeURL(stateJWT, oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(verifier))
	httplib.RenderOK(ctx, w, oapi.GoogleLoginURL{RedirectUrl: redirectURL})
}

func (h GoogleAuthHandler) CallbackFromGoogle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if !h.enabled {
		h.RedirectToResultPage(ctx, w, r, auth.GoogleResultPageNotEnabled)
		return
	}

	state := r.URL.Query().Get("state")

	claims, err := h.authUsecase.VerifyStateJWT(ctx, state)
	if err != nil {
		h.RedirectToResultPage(ctx, w, r, auth.GoogleResultPageInvalidToken)
		return
	}

	id, ok := auth.GetIDFromJWT(claims)
	if !ok {
		h.RedirectToResultPage(ctx, w, r, auth.GoogleResultPageInvalidToken)
		return
	}

	verifier, ok := pkceKeyMap.LoadAndDelete(id)
	if !ok {
		h.RedirectToResultPage(ctx, w, r, auth.GoogleResultPageTryAgain)
		return
	}

	code := r.URL.Query().Get("code")
	token, err := h.conf.Exchange(r.Context(), code, oauth2.VerifierOption(fmt.Sprintf("%v", verifier)))
	if err != nil {
		h.RedirectToResultPage(ctx, w, r, auth.GoogleResultPageInvalidToken)
		return
	}

	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		h.RedirectToResultPage(ctx, w, r, auth.GoogleResultPageInvalidToken)
		return
	}

	jwt := strings.Split(idToken, ".")
	payload := strings.TrimSuffix(jwt[1], "=")
	b, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		h.RedirectToResultPage(ctx, w, r, auth.GoogleResultPageInvalidToken)
		return
	}

	var extraClaims map[string]interface{}
	if err := json.Unmarshal(b, &extraClaims); err != nil {
		h.RedirectToResultPage(ctx, w, r, auth.GoogleResultPageInvalidToken)
		return
	}

	openID, ok := extraClaims["sub"].(string)
	if !ok {
		h.RedirectToResultPage(ctx, w, r, auth.GoogleResultPageInvalidToken)
		return
	}

	loginKey, ok := auth.GetLoginKeyFromJWT(claims)
	if ok {
		h.handleFirstLoginCallback(w, r, openID, loginKey)
	} else {
		redirectTo := auth.GetCurrentPageURLFromJWT(claims)
		h.handleLoginCallback(w, r, openID, redirectTo)
	}
}

func (h GoogleAuthHandler) handleFirstLoginCallback(w http.ResponseWriter, r *http.Request, openID string, loginKey int64) {
	ctx := r.Context()

	userID, err := h.userUsecase.SaveSubByLoginKey(ctx, loginKey, openID)
	if errors.Is(err, user.NotFoundByLoginKeyError{LoginKey: loginKey}) {
		h.RedirectToResultPage(ctx, w, r, auth.GoogleResultPageLoginKeyNotFound)
		return
	} else if err != nil {
		logs.Error(ctx, err)
		h.RedirectToResultPage(ctx, w, r, auth.GoogleResultPageInternalError)
		return
	}

	h.sendTokens(ctx, w, r, userID, "", auditlog.UserActionFirstLogin)
}

func (h GoogleAuthHandler) handleLoginCallback(w http.ResponseWriter, r *http.Request, openID string, redirectTo string) {
	ctx := r.Context()
	userID, err := h.userUsecase.FindUserIDBySub(ctx, openID)
	if errors.Is(err, user.NotFoundBySubError{Sub: openID}) {
		h.RedirectToResultPage(ctx, w, r, auth.GoogleResultPageUserNotFound)
		return
	} else if err != nil {
		logs.Error(ctx, err)
		h.RedirectToResultPage(ctx, w, r, auth.GoogleResultPageInternalError)
		return
	}

	h.sendTokens(ctx, w, r, userID, redirectTo, auditlog.UserActionLogin)
}

func (h GoogleAuthHandler) sendTokens(ctx context.Context, w http.ResponseWriter, r *http.Request, userID user.ID, redirectTo string, action auditlog.UserAction) {
	refreshToken, expiresAt, err := h.authUsecase.CreateRefreshTokens(ctx, userID)
	if err != nil {
		httplib.RenderError(ctx, w, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/auth",
		HttpOnly: true,
		Secure:   true,
		Expires:  expiresAt,
		SameSite: http.SameSiteLaxMode,
	})

	redirect := h.resultPageURL + "?type=" + string(auth.GoogleResultPageSuccess)
	if redirectTo != "" {
		redirect += "&redirectTo=" + url.PathEscape(redirectTo)
	}
	httplib.RenderRedirect(ctx, w, r, redirect)

	ctxlib.SetUserIDForAuditLog(ctx, userID)
	ctxlib.AddAuditLogRecord(ctx, auditlog.UserLogRecord{
		Action:    action,
		Timestamp: time.Now(),
	})
}

func (h GoogleAuthHandler) RedirectToResultPage(ctx context.Context, w http.ResponseWriter, r *http.Request, pageType auth.GoogleResultPageType) {
	redirect := h.resultPageURL + "?type=" + string(pageType)
	httplib.RenderRedirect(ctx, w, r, redirect)
}
