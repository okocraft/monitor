package registry

import (
	"github.com/google/wire"
	"github.com/okocraft/monitor/internal/repositories"
	"github.com/okocraft/monitor/internal/repositories/database"
)

var repositorySet = wire.NewSet(
	repositories.NewAuthRepository,
	repositories.NewUserRepository,
	repositories.NewPermissionRepository,
	repositories.NewAuditLogRepository,
	database.NewTransaction,
)
