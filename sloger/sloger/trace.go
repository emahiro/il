package sloger

import (
	"context"
	"strings"
)

type (
	traceIDKey struct{}
	spanIDKey  struct{}
)

// SetCloudTraceContext sets traceID and spanID to context.
// xcTraceCtx is included in the "X-Cloud-Trace-Context" at a request header.
// xcTraceCtx is retrive from `r.Header.Get("X-Cloud-Trace-Context")`.
func SetCloudTraceContext(ctx context.Context, xcTraceCtx string) context.Context {
	var traceID, spanID string
	tmp := strings.Split(xcTraceCtx, "/")
	if len(tmp) == 2 {
		traceID = tmp[0]
		spanIDStr := strings.Split(tmp[1], ";")
		if len(spanIDStr) == 2 {
			spanID = spanIDStr[0]
		} else {
			spanID = tmp[1]
		}
	}
	ctx = context.WithValue(ctx, traceIDKey{}, traceID)
	ctx = context.WithValue(ctx, spanIDKey{}, spanID)
	return ctx
}
