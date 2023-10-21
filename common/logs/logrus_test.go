package logs

import (
	"context"
	"shield/common/constant"
	"testing"
)

func TestLogrus(t *testing.T) {
	ctx := context.Background()

	CtxInfo(ctx, "test1")

	ctx = context.WithValue(ctx, constant.Trace{}, constant.Trace{
		TraceID: "trace_id",
		SpanID:  "span",
		LogID:   "log_id",
	})

	CtxInfo(ctx, "test2")

	Info("test3")

	CtxInfo(ctx, "123 %s, %+v", "123", 123)
}
