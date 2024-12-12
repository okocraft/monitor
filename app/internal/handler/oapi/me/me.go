package me

import (
	"errors"
	"github.com/okocraft/monitor/internal/domain/user"
	"github.com/okocraft/monitor/internal/usecases"
	"github.com/okocraft/monitor/lib/httplib"
	"net/http"
)

type MeHandler struct {
	usecase usecases.UserUsecase
}

func NewMeHandler(usecase usecases.UserUsecase) MeHandler {
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
