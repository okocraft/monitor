package main

import (
	"context"
	"os"
	"time"

	"github.com/Siroshun09/logs"
	"github.com/okocraft/monitor/internal/config"
	"github.com/okocraft/monitor/internal/handler/logger"
	"github.com/okocraft/monitor/internal/registry"
	"github.com/okocraft/monitor/internal/repositories/database"
	"github.com/okocraft/monitor/lib/errlib"
)

func main() {
	ctx := context.Background()

	debug := os.Getenv("DEBUG") == "true"

	loggerFactory := logger.NewFactory(debug)
	defaultLogger := loggerFactory.NewDefaultLogger()
	ctx = logs.WithContext(ctx, defaultLogger)

	cfg, err := config.NewSetupConfigFromEnv()
	if err != nil {
		logs.Error(ctx, err)
		os.Exit(1)
	}

	db, err := database.New(cfg.DBConfig, 10*time.Minute)
	if err != nil {
		logs.Error(ctx, err)
		os.Exit(1)
	}

	usecase := registry.NewSetupUsecase(db)

	isFresh, err := usecase.IsFreshDatabase(ctx)
	if err != nil {
		logs.Error(ctx, err)
		os.Exit(1)
	}

	if !isFresh && !cfg.ForceSetup {
		logs.Info(ctx, "monitor is already initialized. if you want a login key for admin, please set MONITOR_FORCE_SETUP to true in env")
		os.Exit(1)
	}

	tx := registry.NewTransaction(db)

	var loginKey string
	err = tx.WithTx(ctx, func(ctx context.Context) error {
		role, err := usecase.CreateInitialAdminRole(ctx)
		if err != nil {
			return errlib.AsIs(err)
		}

		user, err := usecase.CreateInitialAdminUser(ctx, role.ID)
		if err != nil {
			return errlib.AsIs(err)
		}

		loginKey, err = usecase.CreateLoginKeyForAdminUser(ctx, user.ID)
		if err != nil {
			return errlib.AsIs(err)
		}

		return nil
	})
	if err != nil {
		logs.Error(ctx, err)
		os.Exit(1)
	}

	logs.Info(ctx, "login key for admin user: "+loginKey)
}
