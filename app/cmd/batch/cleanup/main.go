package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/Siroshun09/logs/v2"
	"github.com/okocraft/monitor/internal/config"
	"github.com/okocraft/monitor/internal/domain/cleanup"
	"github.com/okocraft/monitor/internal/registry"
	"github.com/okocraft/monitor/internal/repositories/database"
)

func main() {
	logger := logs.NewLoggerWithSlog(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	ctx := context.Background()
	ctx = logs.WithContext(ctx, logger)

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
