package kitex

import (
	"context"
	"reflect"
	"shield/common/errs"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

func ErrorHandlerMW(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response interface{}) error {
		if err := next(ctx, request, response); err != nil {
			handleFailedResp(response, err)
			return nil
		}

		handleSuccessResp(response)
		return nil
	}
}

func handleFailedResp(result interface{}, err error) {
	if kRes, ok := result.(interface {
		GetResult() interface{}
		SetSuccess(x interface{})
	}); ok {
		resp := kRes.GetResult()
		if reflect.ValueOf(resp).IsNil() {
			resp = reflect.New(reflect.TypeOf(resp).Elem()).Interface()
			kRes.SetSuccess(resp)
		}

		baseRespField := reflect.ValueOf(resp).Elem().FieldByName("Base")
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

func handleSuccessResp(result interface{}) {
	if kRes, ok := result.(interface {
		GetResult() interface{}
		SetSuccess(x interface{})
	}); ok {
		resp := kRes.GetResult()
		if reflect.ValueOf(resp).IsNil() {
			resp = reflect.New(reflect.TypeOf(resp).Elem()).Interface()
			kRes.SetSuccess(resp)
		}

		baseRespField := reflect.ValueOf(resp).Elem().FieldByName("Base")
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
