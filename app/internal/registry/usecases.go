package registry

import (
	"github.com/google/wire"
	"github.com/okocraft/monitor/internal/usecases"
)

var usecaseSet = wire.NewSet(
	usecases.NewAuthUsecase,
	usecases.NewUserUsecase,
	usecases.NewPermissionUsecase,
	usecases.NewAuditLogUsecase,
)
