package kitex

import (
	"context"
	"shield/common/constant"
	"shield/common/utils"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

func init() {
	idGen = utils.NewIDGenerator(100)
}

var idGen *utils.IDGenerator

func ServerTraceMW(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, args, result interface{}) (err error) {
		var traceID, spanID, logID string
		if gfa, ok := args.(interface{ GetFirstArgument() interface{} }); ok {
			if rv, ok := gfa.GetFirstArgument().(interface{ GetBase() interface{} }); ok {
				if base, ok := rv.GetBase().(interface {
					GetLogID() string
					GetTraceID() string
					GetSpanID() string
				}); ok {
					logID = base.GetLogID()
					traceID = base.GetTraceID()
					spanID = base.GetSpanID()
				}
			}
		}

		if logID == "" {
			logID = idGen.NewLogID()
		}
		if traceID == "" {
			traceID = idGen.NewTraceID()
		}

		ctx = context.WithValue(ctx, constant.Trace{}, constant.Trace{
			LogID:   logID,
			TraceID: traceID,
			SpanID:  idGen.NewSpanID(spanID),
		})

		if err := next(ctx, args, result); err != nil {
			return err
		}

		return nil
	}
}
