package sloger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

var logger *sloger

type sloger struct {
	logger *slog.Logger
	sev    slog.Level
	msg    string
}

func New() {
	logger = &sloger{
		logger: slog.New(slog.NewJSONHandler(os.Stderr, nil)),
	}
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	logger.setMsg(fmt.Sprintf(format, args...))
	logger.setSeverity(slog.LevelInfo)
	logger.output(ctx, defaultCloudLogAttrs(ctx, logger.sev.String(), logger.msg)...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	logger.setMsg(fmt.Sprintf(format, args...))
	logger.setSeverity(slog.LevelError)
	logger.output(ctx, defaultCloudLogAttrs(ctx, logger.sev.String(), logger.msg)...)
}

func (l *sloger) setSeverity(sev slog.Level) {
	l.sev = sev
}

func (l *sloger) setMsg(msg string) {
	l.msg = msg
}

func (l *sloger) output(ctx context.Context, attrs ...slog.Attr) {
	l.logger.LogAttrs(ctx, l.sev, l.msg, attrs...)
}
