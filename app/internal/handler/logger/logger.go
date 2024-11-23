package logger

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/Siroshun09/logs"
	"github.com/Siroshun09/serrors"
)

func NewFactory(debug bool) Factory {
	level := slog.LevelInfo
	if debug {
		level = slog.LevelDebug
	}
	return factory{
		delegate:                slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})),
		includeStackTraceOnWarn: debug,
	}
}

type Factory interface {
	NewLogger(attrFunc func(ctx context.Context) *slog.Attr) logs.Logger
	NewDefaultLogger() logs.Logger
	NewHTTPMiddlewareWithRecover(next http.Handler) http.Handler
}

type factory struct {
	delegate                *slog.Logger
	includeStackTraceOnWarn bool
}

func (f factory) NewLogger(attrFunc func(ctx context.Context) *slog.Attr) logs.Logger {
	return &logger{
		delegate:                f.delegate,
		attrFunc:                attrFunc,
		includeStackTraceOnWarn: f.includeStackTraceOnWarn,
	}
}

func (f factory) NewDefaultLogger() logs.Logger {
	return f.NewLogger(func(ctx context.Context) *slog.Attr {
		return nil
	})
}

type logger struct {
	delegate                *slog.Logger
	attrFunc                func(ctx context.Context) *slog.Attr
	includeStackTraceOnWarn bool
}

func (l logger) Debug(ctx context.Context, msg string) {
	if !l.delegate.Enabled(ctx, slog.LevelDebug) {
		return
	}
	l.delegate.DebugContext(ctx, msg, l.createContextAttr(ctx, nil)...)
}

func (l logger) Info(ctx context.Context, msg string) {
	l.delegate.InfoContext(ctx, msg, l.createContextAttr(ctx, nil)...)
}

func (l logger) Warn(ctx context.Context, err error) {
	if l.includeStackTraceOnWarn {
		stacktrace := slog.String("stacktrace", serrors.GetStackTrace(err).String())
		l.delegate.WarnContext(ctx, err.Error(), l.createContextAttr(ctx, &stacktrace))
	}
	l.delegate.WarnContext(ctx, err.Error(), l.createContextAttr(ctx, nil)...)
}

func (l logger) Error(ctx context.Context, err error) {
	stacktrace := slog.String("stacktrace", serrors.GetStackTrace(err).String())
	l.delegate.ErrorContext(ctx, err.Error(), l.createContextAttr(ctx, &stacktrace)...)
}

func (l logger) createContextAttr(ctx context.Context, stackTraceAttr *slog.Attr) []any {
	attr := l.attrFunc(ctx)
	switch {
	case attr == nil:
		if stackTraceAttr == nil {
			return []any{}
		}
		return []any{*stackTraceAttr}
	case stackTraceAttr == nil:
		return []any{*attr}
	default:
		return []any{*attr, *stackTraceAttr}
	}
}
