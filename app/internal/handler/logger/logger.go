package logger

import (
	"context"
	"log/slog"
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
	NewLogger(attrFunc func(ctx context.Context) []slog.Attr) logs.Logger
	NewDefaultLogger() logs.Logger
}

type factory struct {
	delegate                *slog.Logger
	includeStackTraceOnWarn bool
}

func (f factory) NewLogger(attrFunc func(ctx context.Context) []slog.Attr) logs.Logger {
	return &logger{
		delegate:                f.delegate,
		attrFunc:                attrFunc,
		includeStackTraceOnWarn: f.includeStackTraceOnWarn,
	}
}

func (f factory) NewDefaultLogger() logs.Logger {
	return f.NewLogger(func(ctx context.Context) []slog.Attr {
		return []slog.Attr{}
	})
}

type logger struct {
	delegate                *slog.Logger
	attrFunc                func(ctx context.Context) []slog.Attr
	includeStackTraceOnWarn bool
}

func (l logger) Debug(ctx context.Context, msg string) {
	if !l.delegate.Enabled(ctx, slog.LevelDebug) {
		return
	}
	l.delegate.DebugContext(ctx, msg, l.attrFunc(ctx))
}

func (l logger) Info(ctx context.Context, msg string) {
	l.delegate.InfoContext(ctx, msg, l.attrFunc(ctx))
}

func (l logger) Warn(ctx context.Context, err error) {
	attrs := l.attrFunc(ctx)
	if l.includeStackTraceOnWarn {
		attrs = append(attrs, slog.String("stacktrace", serrors.GetStackTrace(err).String()))
	}
	l.delegate.WarnContext(ctx, err.Error(), attrs)
}

func (l logger) Error(ctx context.Context, err error) {
	attrs := l.attrFunc(ctx)
	attrs = append(attrs, slog.String("stacktrace", serrors.GetStackTrace(err).String()))
	l.delegate.ErrorContext(ctx, err.Error(), attrs)
}
