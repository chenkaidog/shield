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

const baseFieldName = "Base"

type serverBaseReq interface {
	GetLogID() string
	GetTraceID() string
	GetSpanID() string
}


type clientBaseReq interface {
	SetLogID(string)
	SetTraceID(string)
	SetSpanID(string)
}

func ServerTraceMW(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, args, result interface{}) (err error) {
		var traceID, spanID, logID string
		if gfa, ok := args.(interface{ GetFirstArgument() interface{} }); ok {
			if baseReq := getServerBaseReq(gfa.GetFirstArgument()); baseReq != nil {
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

		return next(ctx, args, result)
	}
}

func ClientTraceMW(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, args, result  interface{}) (err error) {
		trace, ok := ctx.Value(constant.Trace{}).(constant.Trace)
		if ok {
			if gfa, ok := args.(interface{ GetFirstArgument() interface{} }); ok {
				if baseReq := getClientBaseReq(gfa.GetFirstArgument()); baseReq != nil {
					baseReq.SetLogID(trace.LogID)
					baseReq.SetTraceID(trace.TraceID)
					baseReq.SetSpanID(trace.SpanID)
				}
			}
		}

		return next(ctx, args, result)
	}
}

func getServerBaseReq(firstArg interface{}) serverBaseReq {
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

	if _, ok := req.Type().FieldByName(baseFieldName); ok {
		if result, ok := req.FieldByName(baseFieldName).Interface().(serverBaseReq); ok {
			return result
		}
		return nil
	}

	return nil
}

func getClientBaseReq(firstArg interface{}) clientBaseReq {
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

	if _, ok := req.Type().FieldByName(baseFieldName); ok {
		if result, ok := req.FieldByName(baseFieldName).Interface().(clientBaseReq); ok {
			return result
		}
		return nil
	}

	return nil
}
