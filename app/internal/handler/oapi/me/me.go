package me

import (
	"errors"
	"github.com/okocraft/monitor/internal/domain/user"
	user2 "github.com/okocraft/monitor/internal/usecases/user"
	"github.com/okocraft/monitor/lib/httplib"
	"net/http"
)

type MeHandler struct {
	usecase user2.UserUsecase
}

func NewMeHandler(usecase user2.UserUsecase) MeHandler {
	return MeHandler{usecase: usecase}
}

func (h MeHandler) GetMe(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	me, err := h.usecase.GetMe(ctx)

	var notFound user.NotFoundByIDError
	if errors.As(err, &notFound) {
		httplib.RenderNotFound(ctx, w, err)
		return
	}
	if err != nil {
		httplib.RenderError(ctx, w, err)
		return
	}

	httplib.RenderOK(ctx, w, me.ToResponse())
}
