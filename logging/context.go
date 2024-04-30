package logging

import (
	"context"
	"log/slog"
)

type key struct {
}

var ctxKey = key{}

func Context(ctx context.Context, loggers ...*slog.Logger) context.Context {
	var logger *slog.Logger
	if len(loggers) > 0 {
		logger = loggers[0]
	} else {
		logger = slog.Default()
	}

	return context.WithValue(ctx, ctxKey, logger)
}

func FromContext(ctx context.Context) *slog.Logger {
	logger, ok := ctx.Value(ctxKey).(*slog.Logger)
	if !ok {
		return slog.Default()
	}
	return logger
}
