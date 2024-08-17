package sloger

import "log/slog"

func WithCloudLogAttr(base *slog.Logger, severity slog.Level) *slog.Logger {
	return base.
		With(slog.String("severity", severity.String()))
}
