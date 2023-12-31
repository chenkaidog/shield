package logs

import (
	"context"
	"shield/common/trace"
	"testing"
)

func TestLogger(t *testing.T) {
	ctx := trace.ContextWithTrace(context.Background(), trace.Trace{
		TraceID: "trace_id",
		SpanID:  "span_id",
		LogID:   "log_id",
	})
	CtxInfof(ctx, "s: %s, int: %d", "abc", 123)
}
