package sloger

import "log/slog"

func WithCloudLogAttr(base *slog.Logger, severity slog.Level) *slog.Logger {
	return base.
		With(slog.String("severity", severity.String()))
}

func defaultCloudLogAttrs(severity, message string) []slog.Attr {
	return []slog.Attr{
		slog.String("severity", severity),
		slog.String("message", message),
	}
}
