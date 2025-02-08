package registry

import (
	"github.com/google/wire"
	"github.com/okocraft/monitor/internal/usecases/auditlog"
	"github.com/okocraft/monitor/internal/usecases/auth"
	"github.com/okocraft/monitor/internal/usecases/permission"
	"github.com/okocraft/monitor/internal/usecases/user"
)

var usecaseSet = wire.NewSet(
	auth.NewAuthUsecase,
	user.NewUserUsecase,
	permission.NewPermissionUsecase,
	auditlog.NewAuditLogUsecase,
)
