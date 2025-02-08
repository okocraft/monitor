package handler

import (
	"github.com/okocraft/monitor/internal/handler/auditlog"
	"github.com/okocraft/monitor/internal/handler/oapi"
	"github.com/okocraft/monitor/internal/handler/oapi/auth"
	"github.com/okocraft/monitor/internal/handler/oapi/me"
	"github.com/okocraft/monitor/internal/handler/oapi/user"
)

var _ oapi.ServerInterface = (*HTTPHandler)(nil)

type HTTPHandler struct {
	auth.AuthHandler
	auth.GoogleAuthHandler
	me.MeHandler
	user.UserHandler

	auditlog.AuditLogMiddleware
}

func NewHTTPHandler(
	authHandler auth.AuthHandler,
	googleAuthHandler auth.GoogleAuthHandler,
	meHandler me.MeHandler,
	userHandler user.UserHandler,
	auditLogMiddleware auditlog.AuditLogMiddleware,
) HTTPHandler {
	return HTTPHandler{
		AuthHandler:        authHandler,
		GoogleAuthHandler:  googleAuthHandler,
		MeHandler:          meHandler,
		UserHandler:        userHandler,
		AuditLogMiddleware: auditLogMiddleware,
	}
}
