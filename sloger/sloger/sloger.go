package sloger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

var logger *slog.Logger

func New() {
	logger = WithCloudLogAttr(slog.New(slog.NewJSONHandler(os.Stderr, nil)), slog.LevelInfo)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	logger.LogAttrs(ctx, slog.LevelInfo, msg)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	logger.LogAttrs(ctx, slog.LevelError, msg)
}
