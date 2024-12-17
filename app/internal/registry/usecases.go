package registry

import (
	"github.com/google/wire"
	"github.com/okocraft/monitor/internal/usecases"
)

var usecaseSet = wire.NewSet(
	usecases.NewUserUsecase,
	usecases.NewAuthUsecase,
	usecases.NewAuditLogUsecase,
)
