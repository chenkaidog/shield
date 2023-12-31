package kitex

import (
	"context"
	"reflect"
	"shield/common/errs"
	"shield/common/logs"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

func ServerErrorHandlerMW(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response interface{}) error {
		err := next(ctx, request, response)

		if kRes, ok := response.(interface {
			GetResult() interface{}
			SetSuccess(x interface{})
		}); ok {
			resp := kRes.GetResult()
			if val := reflect.ValueOf(resp); val.Kind() == reflect.Ptr && val.IsNil() {
				resp = reflect.New(val.Type().Elem()).Interface()
				kRes.SetSuccess(resp)
			}

			if err != nil {
				logs.CtxDebugf(ctx, "err type: %T", err)
				return handleServerFailedResp(resp, err)
			}

			handleServerSuccessResp(resp)
			return nil
		}

		return err
	}
}

func handleServerFailedResp(resp interface{}, err error) error {
	baseRespField := reflect.ValueOf(resp).Elem().FieldByName(baseFieldName)
	if baseRespField.Kind() == reflect.Ptr && baseRespField.IsNil() {
		logs.Debug("base is nil")
		baseRespField.Set(reflect.New(baseRespField.Type().Elem()))
	}

	if baseInf, ok := baseRespField.Interface().(interface {
		SetCode(val int32)
		SetMsg(val string)
		SetSuccess(val bool)
	}); ok {
		logs.Debugf("interface assert true")
		baseInf.SetSuccess(false)
		if bizErr, ok := convert2BizErr(err); ok {
			logs.Debugf("is biz err")
			baseInf.SetCode(bizErr.Code())
			baseInf.SetMsg(bizErr.Msg())
			return nil
		}
	}

	return err
}

func handleServerSuccessResp(resp interface{}) {
	baseRespField := reflect.ValueOf(resp).Elem().FieldByName(baseFieldName)
	if baseRespField.Kind() == reflect.Ptr && baseRespField.IsNil() {
		baseRespField.Set(reflect.New(baseRespField.Type().Elem()))
	}

	if baseInf, ok := baseRespField.Interface().(interface {
		SetCode(val int32)
		SetMsg(val string)
		SetSuccess(val bool)
	}); ok {
		baseInf.SetSuccess(true)
		baseInf.SetCode(errs.Success.Code())
		baseInf.SetMsg(errs.Success.Msg())
	}
}

func ClientErrorHandlerMW(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response interface{}) error {
		if err := next(ctx, request, response); err != nil {

			// error occurs in middleware
			if bizErr, ok := convert2BizErr(err); ok {
				return bizErr
			}

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

func convert2BizErr(err error) (errs.Error, bool) {
	if err == nil {
		return nil, false
	}

	if bizErr, ok := err.(errs.Error); ok {
		return bizErr, true
	}

	if detailErr, ok := err.(*kerrors.DetailedError); ok {
		if bizErr, ok := detailErr.Unwrap().(errs.Error); ok {
			return bizErr, true
		}
	}

	return nil, false
}
