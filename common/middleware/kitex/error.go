package kitex

import (
	"context"
	"reflect"
	"shield/common/errs"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

func ServerErrorHandlerMW(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response interface{}) error {
		if err := next(ctx, request, response); err != nil {
			handleServerFailedResp(response, err)
			return nil
		}

		handleServerSuccessResp(response)
		return nil
	}
}

func ClientErrorHandlerMW(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response interface{}) error {
		if err := next(ctx, request, response); err != nil {
			// RPC发生异常，直接返回RPC error
			return errs.RpcError.SetErr(err)
		}

		// rpc正常，检测是否有业务异常
		if kRes, ok := response.(interface{ GetResult() interface{} }); ok {
			baseRespField := reflect.ValueOf(kRes.GetResult()).Elem().FieldByName(baseFieldName)
			if baseResp, ok := baseRespField.Interface().(interface {
				GetCode() int32
				GetMsg() string
			}); ok {
				if baseResp.GetCode() != errs.Success.Code() {
					return errs.New(baseResp.GetCode(), baseResp.GetMsg())
				}
			}
		}

		return nil
	}
}

func handleServerFailedResp(result interface{}, err error) {
	if kRes, ok := result.(interface {
		GetResult() interface{}
		SetSuccess(x interface{})
	}); ok {
		resp := kRes.GetResult()
		if reflect.ValueOf(resp).IsNil() {
			resp = reflect.New(reflect.TypeOf(resp).Elem()).Interface()
			kRes.SetSuccess(resp)
		}

		baseRespField := reflect.ValueOf(resp).Elem().FieldByName(baseFieldName)
		baseResp := reflect.New(baseRespField.Type().Elem())
		baseRespField.Set(baseResp)
		if baseInf, ok := baseResp.Interface().(interface {
			SetCode(val int32)
			SetMsg(val string)
			SetSuccess(val bool)
		}); ok {
			baseInf.SetSuccess(false)
			if bizErr, ok := err.(errs.Error); ok {
				baseInf.SetCode(bizErr.Code())
				baseInf.SetMsg(bizErr.Msg())
			} else {
				baseInf.SetCode(errs.ServerError.Code())
				baseInf.SetMsg(err.Error())
			}
		}
	}
}

func handleServerSuccessResp(result interface{}) {
	if kRes, ok := result.(interface {
		GetResult() interface{}
		SetSuccess(x interface{})
	}); ok {
		resp := kRes.GetResult()
		if reflect.ValueOf(resp).IsNil() {
			resp = reflect.New(reflect.TypeOf(resp).Elem()).Interface()
			kRes.SetSuccess(resp)
		}

		baseRespField := reflect.ValueOf(resp).Elem().FieldByName(baseFieldName)
		baseResp := reflect.New(baseRespField.Type().Elem())
		baseRespField.Set(baseResp)
		if baseInf, ok := baseResp.Interface().(interface {
			SetCode(val int32)
			SetMsg(val string)
			SetSuccess(val bool)
		}); ok {
			baseInf.SetSuccess(true)
			baseInf.SetCode(errs.Success.Code())
			baseInf.SetMsg(errs.Success.Msg())
		}
	}
}
