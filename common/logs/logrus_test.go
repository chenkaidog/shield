package logs

import (
	"context"
	"os"
	"shield/common/constant"
	"testing"
)

func TestLogrus(t *testing.T) {
	ctx := context.Background()

	CtxInfo(ctx, "test1")

	ctx = context.WithValue(ctx, constant.Trace{}, constant.Trace{
		TraceID: "trace_id",
		SpanID:  "span",
		PSpanID: "psan",
		LogID:   "log_id",
	})

	CtxInfo(ctx, "test2")

	Info("test3")

	t.Log(os.Getwd())
}
