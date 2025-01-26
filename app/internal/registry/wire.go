//go:build wireinject

package registry

import (
	"github.com/google/wire"
	"github.com/okocraft/monitor/internal/config"
	"github.com/okocraft/monitor/internal/handler"
	"github.com/okocraft/monitor/internal/repositories/database"
	"github.com/okocraft/monitor/internal/usecases"
)

//go:generate go run github.com/google/wire/cmd/wire@v0.6.0

func NewHTTPHandler(cfg config.HTTPServerConfig, db database.DB) (handler.HTTPHandler, error) {
	wire.Build(
		configSet,
		repositorySet,
		usecaseSet,
		handlerSet,
		handler.NewHTTPHandler,
	)
	return handler.HTTPHandler{}, nil
}

func NewCleanupUsecase(db database.DB) usecases.CleanupUsecase {
	wire.Build(
		repositorySet,
		usecases.NewCleanupUsecase,
	)
	return usecases.NewCleanupUsecase(nil, nil)
}
