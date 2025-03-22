package logger

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/Siroshun09/logs"
	"github.com/Siroshun09/serrors"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/okocraft/monitor/lib/ctxlib"
)

func (f factory) NewHTTPMiddlewareWithRecover(next http.Handler) http.Handler {
	l := f.NewLogger(func(ctx context.Context) *slog.Attr {
		attr := ctxlib.GetHTTPAccessLog(ctx).ToAttr()
		return &attr
	})

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				l.Error(r.Context(), serrors.Errorf("%v", rvr))
			}
		}()

		ctx := r.Context()
		ctx = logs.WithContext(r.Context(), l)
		ctx, accessLog := ctxlib.InitHTTPAccessLog(ctx)

		accessLog.FromRequest(r)

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		r = r.WithContext(ctx)
		next.ServeHTTP(ww, r)

		if accessLog.Response != nil {
			accessLog.Response.ResponseSize = ww.BytesWritten()
			accessLog.Response.FinishedAt = time.Now()
		}

		switch {
		case accessLog.Response == nil:
			l.Info(ctx, "http access")
		case accessLog.Response.Error == nil:
			l.Info(ctx, "http access handled")
		case accessLog.Response.StatusCode < http.StatusInternalServerError:
			l.Warn(ctx, accessLog.Response.Error)
		default:
			l.Error(ctx, accessLog.Response.Error)
		}
	})
}
