package kitex

import (
	"context"
	"reflect"
	"shield/common/constant"
	"shield/common/utils/idgen"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

func init() {
	idGen = idgen.NewIDGenerator(100)
}

var idGen *idgen.IDGenerator

type BaseReq interface {
    GetLogID() string
    GetTraceID() string
    GetSpanID() string
}

func ServerTraceMW(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, args, result interface{}) (err error) {
		var traceID, spanID, logID string
		if gfa, ok := args.(interface{ GetFirstArgument() interface{} }); ok {
	        if baseReq := getBaseReq(gfa.GetFirstArgument()); baseReq != nil {
	            logID = baseReq.GetLogID()
                traceID = baseReq.GetTraceID()
                spanID = baseReq.GetSpanID()
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

func getBaseReq(firstArg interface{}) BaseReq {
	req := reflect.ValueOf(firstArg)
	if req.Kind() == reflect.Ptr {
	    if req.IsNil() {
            return nil
        }
        req = req.Elem()
	}
	if req.Kind() != reflect.Struct {
	    return nil
	}

	if _, ok := req.Type().FieldByName("Base"); ok {
	    if result, ok :=  req.FieldByName("Base").Interface().(BaseReq); ok {
			return result
		}
		return nil
	}

	return nil
}