package main

import (
	"context"
	"errors"
	"github.com/Siroshun09/serrors"
	"github.com/go-chi/cors"
	nethttpmiddleware "github.com/oapi-codegen/nethttp-middleware"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Siroshun09/logs"
	"github.com/go-chi/chi/v5"
	"github.com/okocraft/monitor/internal/config"
	"github.com/okocraft/monitor/internal/handler/logger"
	"github.com/okocraft/monitor/internal/handler/oapi"
	"github.com/okocraft/monitor/internal/registry"
)

func main() {
	ctx := context.Background()

	debug := os.Getenv("DEBUG") == "true"

	loggerFactory := logger.NewFactory(debug)
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

	server, err := createHTTPServer(loggerFactory, httpHandler, cfg, debug)
	if err != nil {
		defaultLogger.Error(ctx, err)
		os.Exit(1)
	}

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGTERM, os.Interrupt, os.Kill)
	defer stop()

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

func createHTTPServer(loggerFactory logger.Factory, httpHandler oapi.HTTPHandler, cfg config.HTTPServerConfig, printUnknownOrigin bool) (http.Server, error) {
	r := chi.NewRouter()
	r.Use(loggerFactory.NewHTTPMiddlewareWithRecover)
	r.Use(cors.Handler(cors.Options{
		AllowOriginFunc: func(r *http.Request, origin string) bool {
			if _, ok := cfg.AllowedOrigins[origin]; ok {
				return true
			}

			if printUnknownOrigin {
				logs.Info(r.Context(), "Unknown origin: "+origin)
			}
			return false
		},
	}))

	swagger, err := oapi.GetSwagger()
	if err != nil {
		return http.Server{}, serrors.WithStackTrace(err)
	}

	swagger.Servers = nil
	r.Use(nethttpmiddleware.OapiRequestValidator(swagger))

	return http.Server{
		Addr:    ":" + cfg.Port,
		Handler: oapi.HandlerFromMux(httpHandler, r),
	}, nil
}
