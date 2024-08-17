package sloger

import (
	"context"
	"log/slog"
	"os"
)

func WithCloudLogAttr(base *slog.Logger, severity slog.Level) *slog.Logger {
	return base.
		With(slog.String("severity", severity.String()))
}

func defaultCloudLogAttrs(ctx context.Context, severity, message string) []slog.Attr {
	traceID, ok := ctx.Value(traceIDKey{}).(string)
	if !ok {
		traceID = ""
	}
	spanID, ok := ctx.Value(spanIDKey{}).(string)
	if !ok {
		spanID = ""
	}
	return []slog.Attr{
		slog.String("severity", severity),
		slog.String("message", message),
		slog.String("logging.googleapis.com/spanId", spanID),
		slog.String("logging.googleapis.com/trace", "projects/"+os.Getenv("GCP_PROJECT_ID")+"/traces/"+traceID),
	}
}
