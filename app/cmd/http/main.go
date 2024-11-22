package main

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Siroshun09/logs"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/okocraft/monitor/internal/config"
	"github.com/okocraft/monitor/internal/handler/logger"
	"github.com/okocraft/monitor/internal/handler/oapi"
	"github.com/okocraft/monitor/internal/registry"
)

func main() {
	ctx := context.Background()

	loggerFactory := logger.NewFactory(os.Getenv("DEBUG") == "true")
	defaultLogger := loggerFactory.NewDefaultLogger()
	ctx = logs.WithContext(ctx, defaultLogger)

	cfg, err := config.NewHTTPServerConfigFromEnv()
	if err != nil {
		defaultLogger.Error(ctx, err)
		os.Exit(1)
	}

	httpHandler, err := registry.NewHTTPHandler()
	if err != nil {
		defaultLogger.Error(ctx, err)
		os.Exit(1)
	}

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGTERM, os.Interrupt, os.Kill)
	defer stop()

	server := createHTTPServer(httpHandler, cfg.Port)
	go func() {
		if srvErr := server.ListenAndServe(); srvErr != nil {
			if errors.Is(srvErr, http.ErrServerClosed) {
				defaultLogger.Info(ctx, "http server closed")
			} else {
				defaultLogger.Error(ctx, err)
			}
		}
	}()

	defaultLogger.Info(ctx, "monitor-app-http has been started with port "+cfg.Port)

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		defaultLogger.Error(ctx, err)
		os.Exit(1)
	}

	defaultLogger.Info(ctx, "monitor-app-http has been stopped")
}

func createHTTPServer(httpHandler oapi.HTTPHandler, port string) http.Server {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(
		middleware.Recoverer,
	)

	return http.Server{
		Addr:    ":" + port,
		Handler: oapi.HandlerFromMux(httpHandler, r),
	}
}
