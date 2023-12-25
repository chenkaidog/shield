package hertz

import (
	"context"
	"shield/common/trace"
	"shield/common/utils/idgen"

	"github.com/cloudwego/hertz/pkg/app"
)

func init() {
	idGen = idgen.NewIDGenerator(100)
}

var idGen *idgen.IDGenerator

const (
	headerKeyTraceId = "X-Trace-ID"
	headerKeyLogId   = "X-Log-ID"
	headerKeySpanId  = "X-Span-ID"
)

func ServerTraceMW() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		traceID := c.Request.Header.Get(headerKeyTraceId)
		if traceID == "" {
			traceID = idGen.NewTraceID()
		}

		logID := c.Request.Header.Get(headerKeyLogId)
		if logID == "" {
			logID = idGen.NewLogID()
		}

		pspanID := c.Request.Header.Get(headerKeySpanId)
		spanID := idGen.NewSpanID(pspanID)

		ctx = trace.ContextWithTrace(ctx, trace.Trace{
			LogID:   logID,
			TraceID: traceID,
			SpanID:  spanID,
		})

		c.Header(headerKeyTraceId, traceID)
		c.Header(headerKeyLogId, logID)
		c.Header(headerKeySpanId, spanID)

		c.Next(ctx)
	}
}
