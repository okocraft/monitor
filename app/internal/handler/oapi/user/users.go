package user

import (
	"net/http"

	"github.com/Siroshun09/serrors"
	"github.com/okocraft/monitor/internal/domain/permission"
	"github.com/okocraft/monitor/internal/domain/user"
	"github.com/okocraft/monitor/internal/handler/oapi"
	permissionUsecase "github.com/okocraft/monitor/internal/usecases/permission"
	userUsecase "github.com/okocraft/monitor/internal/usecases/user"
	"github.com/okocraft/monitor/lib/ctxlib"
	"github.com/okocraft/monitor/lib/httplib"
)

type UserHandler struct {
	usecase           userUsecase.UserUsecase
	permissionUsecase permissionUsecase.PermissionUsecase
}

func NewUserHandler(usecase userUsecase.UserUsecase, permissionUsecase permissionUsecase.PermissionUsecase) UserHandler {
	return UserHandler{
		usecase:           usecase,
		permissionUsecase: permissionUsecase,
	}
}

func (h UserHandler) GetUsersByIds(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, ok := ctxlib.GetUserID(ctx)
	if !ok {
		httplib.RenderOK(ctx, w, []oapi.UUID{})
		return
	}

	if hasPermission, err := h.permissionUsecase.HasPermission(ctx, userID, permission.UserList); err != nil {
		httplib.RenderError(ctx, w, err)
		return
	} else if !hasPermission {
		httplib.RenderOK(ctx, w, []oapi.UUID{})
		return
	}

	req, err := httplib.DecodeBody[oapi.GetUsersByIdsJSONRequestBody](r)
	if err != nil {
		httplib.RenderBadRequest(ctx, w, err)
		return
	}

	users, err := h.usecase.GetUsersWithRoleByUUIDs(ctx, req)
	if err != nil {
		httplib.RenderError(ctx, w, err)
		return
	}

	if len(users) != len(req) {
		httplib.RenderNotFound(ctx, w, serrors.New("some users not found"))
		return
	}

	res := make([]oapi.User, len(users))
	for i, u := range users {
		res[i] = u.ToResponse()
	}

	httplib.RenderOK(ctx, w, res)
}

func (h UserHandler) SearchUsers(w http.ResponseWriter, r *http.Request, params oapi.SearchUsersParams) {
	ctx := r.Context()

	userID, ok := ctxlib.GetUserID(ctx)
	if !ok {
		httplib.RenderOK(ctx, w, []oapi.UUID{})
		return
	}

	if hasPermission, err := h.permissionUsecase.HasPermission(ctx, userID, permission.UserList); err != nil {
		httplib.RenderError(ctx, w, err)
		return
	} else if !hasPermission {
		httplib.RenderOK(ctx, w, []oapi.UUID{})
		return
	}

	searchParams, err := user.NewSearchParamsFromRequest(params)
	if err != nil {
		httplib.RenderBadRequest(ctx, w, err)
		return
	}

	ids, err := h.usecase.SearchForUserUUIDs(ctx, searchParams)
	if err != nil {
		httplib.RenderError(ctx, w, err)
		return
	}

	httplib.RenderOK(ctx, w, ids)
}
