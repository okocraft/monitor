package user

import (
	"github.com/Siroshun09/serrors"
	"github.com/okocraft/monitor/internal/handler/oapi"
	"github.com/okocraft/monitor/internal/usecases"
	"github.com/okocraft/monitor/lib/httplib"
	"net/http"
)

type UserHandler struct {
	usecase usecases.UserUsecase
}

func NewUserHandler(usecase usecases.UserUsecase) UserHandler {
	return UserHandler{
		usecase: usecase,
	}
}

func (h UserHandler) GetUsersByIds(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

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
