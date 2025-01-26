package main

import (
	"context"
	"fmt"
	"github.com/Siroshun09/logs"
	"github.com/okocraft/monitor/internal/config"
	"github.com/okocraft/monitor/internal/domain/cleanup"
	"github.com/okocraft/monitor/internal/handler/logger"
	"github.com/okocraft/monitor/internal/registry"
	"github.com/okocraft/monitor/internal/repositories/database"
	"os"
	"time"
)

func main() {
	ctx := context.Background()

	debug := os.Getenv("DEBUG") == "true"

	loggerFactory := logger.NewFactory(debug)
	defaultLogger := loggerFactory.NewDefaultLogger()
	ctx = logs.WithContext(ctx, defaultLogger)

	cfg, err := config.NewCleanupConfigFromEnv()
	if err != nil {
		logs.Error(ctx, err)
		os.Exit(1)
	}

	db, err := database.New(cfg.DBConfig, 10*time.Minute)
	if err != nil {
		logs.Error(ctx, err)
		os.Exit(1)
	}

	usecase := registry.NewCleanupUsecase(db)

	result, err := usecase.CleanupExpiredTokens(ctx, cleanup.Param{
		AccessTokenExpiredAt:  time.Now().Add(-cfg.AccessTokenExpireDuration),
		RefreshTokenExpiredAt: time.Now().Add(-cfg.RefreshTokenExpireDuration),
	})
	if err != nil {
		logs.Error(ctx, err)
		os.Exit(1)
	}

	logs.Info(ctx, fmt.Sprintf("%d access tokens and %d refresh tokens have been deleted", result.AccessTokens, result.RefreshTokens))
}
