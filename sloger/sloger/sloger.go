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
	sev := slog.LevelInfo
	output(ctx, sev, msg, defaultCloudLogAttrs(sev.String(), msg)...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	sev := slog.LevelError
	output(ctx, sev, msg, defaultCloudLogAttrs(sev.String(), msg)...)
}

func output(ctx context.Context, severity slog.Level, msg string, attrs ...slog.Attr) {
	logger.LogAttrs(ctx, severity, msg, attrs...)
}
