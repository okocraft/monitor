package registry

import (
	"github.com/google/wire"
	"github.com/okocraft/monitor/internal/handler/auditlog"
	"github.com/okocraft/monitor/internal/handler/oapi/auth"
	"github.com/okocraft/monitor/internal/handler/oapi/me"
	"github.com/okocraft/monitor/internal/handler/oapi/role"
	"github.com/okocraft/monitor/internal/handler/oapi/user"
)

var handlerSet = wire.NewSet(
	auth.NewAuthHandler,
	auth.NewGoogleAuthHandler,
	me.NewMeHandler,
	role.NewRoleHandler,
	user.NewUserHandler,
	auditlog.NewAuditLogMiddleware,
)
