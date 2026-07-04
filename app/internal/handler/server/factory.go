package server

import (
	"context"
	"log/slog"
	"maps"
	"net/http"
	"slices"
	"time"

	"github.com/Siroshun09/go-httplib"
	"github.com/Siroshun09/go-httplib/httplog"
	"github.com/Siroshun09/go-httplib/runner"
	"github.com/Siroshun09/logs/errorlogs/v2"
	"github.com/Siroshun09/logs/v2"
	"github.com/Siroshun09/serrors/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/okocraft/monitor/internal/config"
	"github.com/okocraft/monitor/internal/handler"
	"github.com/okocraft/monitor/internal/handler/oapi"
)

type HTTPServerFactory struct {
	cfg     config.HTTPServerConfig
	logger  logs.Logger
	handler handler.HTTPHandler
	debug   bool
}

func NewHTTPServerFactory(cfg config.HTTPServerConfig, logger logs.Logger, handler handler.HTTPHandler, debug bool) HTTPServerFactory {
	return HTTPServerFactory{
		cfg: cfg,
		logger: errorlogs.NewLogger(
			httplog.NewHTTPAttrLogger(logger),
			&errorlogs.Option{
				PrintStackTraceOnWarn:               debug,
				PrintCurrentStackTraceIfNotAttached: true,
			},
		),
		handler: handler,
		debug:   debug,
	}
}

func (f HTTPServerFactory) NewHTTPServer() runner.HTTPServerRunner {
	r := chi.NewRouter()

	r.Use(f.newRecoverer)
	r.Use(f.newLoggerMiddleware)
	r.Use(f.newRecovererForAPIHandler)
	r.Use(cors.Handler(cors.Options{
		AllowOriginFunc: func(r *http.Request, origin string) bool {
			if _, ok := f.cfg.AllowedOrigins[origin]; ok {
				return true
			}

			if f.debug {
				logs.Warn(r.Context(), serrors.New("unknown origin", slog.String("origin", origin)))
			}
			return false
		},
		AllowedOrigins:   slices.Collect(maps.Keys(f.cfg.AllowedOrigins)),
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Use(
		f.handler.AuthHandler.NewAuthMiddleware,
		f.handler.AuditLogMiddleware.NewHTTPMiddleware,
	)

	return runner.NewHTTPServerRunner(
		&http.Server{
			Addr:    ":" + f.cfg.Port,
			Handler: oapi.HandlerFromMux(f.handler, r),
		},
		func(ctx context.Context, err error) {
			logs.Error(ctx, err)
		},
		func(ctx context.Context, rvr any) {
			logs.Error(ctx, serrors.New("panic occurred", slog.Any("panic", rvr)))
		},
	)
}

func (f HTTPServerFactory) newRecoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				logs.Error(r.Context(), serrors.New("panic occurred", slog.Any("panic", rvr)))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (f HTTPServerFactory) newRecovererForAPIHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				httplib.RenderInternalServerError(r.Context(), w, serrors.New("panic occurred", slog.Any("panic", rvr)))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (f HTTPServerFactory) newLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = logs.WithContext(ctx, f.logger)

		requestLog := httplib.NewRequestLog(r, time.Now())
		ctx = httplib.WithRequestLog(ctx, requestLog)

		responseLog := httplib.ResponseLog{}
		ctx = httplib.WithResponseLogPtr(ctx, &responseLog)

		next.ServeHTTP(w, r.WithContext(ctx))

		latency := time.Now().Sub(requestLog.Timestamp)
		ctx = httplib.WithLatency(ctx, latency)

		switch {
		case responseLog.Error == nil:
			logs.Info(ctx, "http access handled")
		case responseLog.StatusCode < http.StatusInternalServerError:
			logs.Warn(ctx, responseLog.Error)
		default:
			logs.Error(ctx, responseLog.Error)
		}
	})
}
