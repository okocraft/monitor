package registry

import (
	"github.com/google/wire"
	"github.com/okocraft/monitor/internal/repositories/auditlog"
	"github.com/okocraft/monitor/internal/repositories/auth"
	"github.com/okocraft/monitor/internal/repositories/database"
	"github.com/okocraft/monitor/internal/repositories/permission"
	"github.com/okocraft/monitor/internal/repositories/user"
)

var repositorySet = wire.NewSet(
	auth.NewAuthRepository,
	user.NewUserRepository,
	permission.NewPermissionRepository,
	auditlog.NewAuditLogRepository,
	database.NewTransaction,
)
