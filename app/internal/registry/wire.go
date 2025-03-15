//go:build wireinject

package registry

import (
	"github.com/google/wire"
	"github.com/okocraft/monitor/internal/config"
	"github.com/okocraft/monitor/internal/handler"
	"github.com/okocraft/monitor/internal/repositories/database"
	"github.com/okocraft/monitor/internal/usecases/cleanup"
	"github.com/okocraft/monitor/internal/usecases/setup"
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

func NewCleanupUsecase(db database.DB) cleanup.CleanupUsecase {
	wire.Build(
		repositorySet,
		cleanup.NewCleanupUsecase,
	)
	return cleanup.NewCleanupUsecase(nil, nil)
}

func NewSetupUsecase(db database.DB) setup.SetupUsecase {
	wire.Build(
		repositorySet,
		setup.NewSetupUsecase,
	)
	return setup.NewSetupUsecase(nil, nil, nil)
}

func NewTransaction(db database.DB) database.Transaction {
	wire.Build(
		database.NewTransaction,
	)
	return database.NewTransaction(nil)
}
