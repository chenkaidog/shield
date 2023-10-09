package kitex

import (
	"context"
	"reflect"
	"shield/account/kitex_gen/base"
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
	baseResp := base.NewBaseResp()
	baseResp.SetSuccess(false)
	if bizErr, ok := err.(errs.Error); ok {
		baseResp.SetCode(bizErr.Code())
		baseResp.SetMsg(bizErr.Msg())
	} else {
		baseResp.SetCode(errs.ServerError.Code())
		baseResp.SetMsg(err.Error())
	}

	if kRes, ok := result.(interface {
		GetResult() interface{}
		SetSuccess(x interface{})
	}); ok {
		resp := kRes.GetResult()
		if reflect.ValueOf(resp).IsNil() {
			resp = reflect.New(reflect.TypeOf(resp).Elem()).Interface()
			kRes.SetSuccess(resp)
		}

		if rv, ok := resp.(interface {
			SetBase(val *base.BaseResp)
		}); ok {
			rv.SetBase(baseResp)
		}
	}
}

func handleSuccessResp(result interface{}) {
	baseResp := base.NewBaseResp()
	baseResp.SetCode(errs.Success.Code())
	baseResp.SetMsg(errs.Success.Msg())
	baseResp.SetSuccess(true)

	if kRes, ok := result.(interface {
		GetResult() interface{}
		SetSuccess(x interface{})
	}); ok {
		resp := kRes.GetResult()
		if reflect.ValueOf(resp).IsNil() {
			resp = reflect.New(reflect.TypeOf(resp).Elem()).Interface()
			kRes.SetSuccess(resp)
		}

		if rv, ok := resp.(interface {
			SetBase(val *base.BaseResp)
		}); ok {
			rv.SetBase(baseResp)
		}
	}
}
