package handler

import (
	"github.com/okocraft/monitor/internal/handler/oapi"
	"github.com/okocraft/monitor/internal/handler/oapi/auth"
	"github.com/okocraft/monitor/internal/handler/oapi/me"
)

var _ oapi.ServerInterface = (*HTTPHandler)(nil)

type HTTPHandler struct {
	auth.AuthHandler
	auth.GoogleAuthHandler
	me.MeHandler
}

func NewHTTPHandler(
	authHandler auth.AuthHandler,
	googleAuthHandler auth.GoogleAuthHandler,
	meHandler me.MeHandler,
) HTTPHandler {
	return HTTPHandler{
		AuthHandler:       authHandler,
		GoogleAuthHandler: googleAuthHandler,
		MeHandler:         meHandler,
	}
}
