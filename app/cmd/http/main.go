package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/okocraft/monitor/internal/handler/server"
	"github.com/okocraft/monitor/internal/repositories/database"

	"github.com/Siroshun09/logs/v2"
	"github.com/okocraft/monitor/internal/config"
	"github.com/okocraft/monitor/internal/registry"
)

func main() {
	logger := logs.NewLoggerWithSlog(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	ctx := context.Background()
	ctx = logs.WithContext(ctx, logger)

	debug := os.Getenv("DEBUG") == "true"

	cfg, err := config.NewHTTPServerConfigFromEnv()
	if err != nil {
		logger.Error(ctx, err)
		os.Exit(1)
	}

	db, err := database.New(cfg.DBConfig, 10*time.Minute)
	if err != nil {
		logger.Error(ctx, err)
		os.Exit(1)
	}

	httpHandler, err := registry.NewHTTPHandler(cfg, db)
	if err != nil {
		logger.Error(ctx, err)
		os.Exit(1)
	}

	httpServer := server.NewHTTPServerFactory(cfg, logger, httpHandler, debug).NewHTTPServer()

	srvCtx, stop := httpServer.Run(ctx)
	logger.Info(ctx, "http server started")
	defer stop()

	<-srvCtx.Done()
	if err := httpServer.Shutdown(1 * time.Minute); err != nil {
		logger.Error(ctx, err)
		os.Exit(1)
	}

	logger.Info(ctx, "http server has been stopped")
	os.Exit(0)
}
