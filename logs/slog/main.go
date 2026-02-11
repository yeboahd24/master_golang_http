package main

import (
	"context"
	"log/slog"
	"os"
	"time"
)

type CustomHandler struct {
	h slog.Handler
}

func (c *CustomHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return c.h.Enabled(ctx, level)
}

func (c *CustomHandler) Handle(ctx context.Context, r slog.Record) error {
	r.Time = time.Now()
	return c.h.Handle(ctx, r)
}

func (c *CustomHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &CustomHandler{h: c.h.WithAttrs(attrs)}
}

func (c *CustomHandler) WithGroup(group string) slog.Handler {
	return &CustomHandler{h: c.h.WithGroup(group)}
}

func main() {
	slog.Info("hello world")
	slog.Info("hello world", "user", os.Getenv("USER")) // key-value pair

	// Everything turns into a key-value pair
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("hello world")
	logger.Info("hello world", "user", os.Getenv("USER"))

	// JsonHandler
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("hello world")
	logger.Info("hello world", "user", os.Getenv("USER"))

	// JsonHandler with options
	logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	logger.Info("hello world")
	logger.Info("hello world", "user", os.Getenv("USER"))

	// CustomHandler
	handler := slog.NewJSONHandler(os.Stdout, nil)
	customHandler := &CustomHandler{h: handler}
	logger = slog.New(customHandler)
	logger.Info("hello world")
	logger.Info("hello world", "user", os.Getenv("USER"))
}
