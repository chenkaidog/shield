package kitex

import (
	"reflect"
	"shield/common/errs"
	"testing"

	"github.com/stretchr/testify/assert"
)

type BaseResp struct {
	Success bool              `thrift:"success,1" frugal:"1,default,bool" json:"success"`
	Code    int32             `thrift:"code,2" frugal:"2,default,i32" json:"code"`
	Msg     string            `thrift:"msg,3" frugal:"3,default,string" json:"msg"`
	Extra   map[string]string `thrift:"extra,255" frugal:"255,default,map<string:string>" json:"extra"`
}

func (base *BaseResp) SetCode(val int32) {
	base.Code = val
}

func (base *BaseResp) SetMsg(val string) {
	base.Msg = val
}

func (base *BaseResp) SetSuccess(val bool) {
	base.Success = val
}

type resp struct {
	Base *BaseResp
}

type result struct {
	Success *resp
}

func (res *result) GetResult() interface{} {
	return res.Success
}

func (res *result) SetSuccess(x interface{}) {
	res.Success = x.(*resp)
}

func Test_handleServerSuccessResp(t *testing.T) {
	ret := new(result)

	var v any = ret
	if kRes, ok := v.(interface {
		GetResult() interface{}
		SetSuccess(x interface{})
	}); ok {
		resp := kRes.GetResult()
		if val := reflect.ValueOf(resp); val.Kind() == reflect.Ptr && val.IsNil() {
			resp = reflect.New(val.Type().Elem()).Interface()
			kRes.SetSuccess(resp)
		}
		handleServerSuccessResp(kRes.GetResult())
	}

	assert.True(t, ret.Success.Base.Success)
	assert.Equal(t, errs.Success.Code(), ret.Success.Base.Code)
	assert.Equal(t, errs.Success.Msg(), ret.Success.Base.Msg)
}

func Test_handleServerFailedResp(t *testing.T) {
	ret := new(result)

	var v any = ret
	if kRes, ok := v.(interface {
		GetResult() interface{}
		SetSuccess(x interface{})
	}); ok {
		resp := kRes.GetResult()
		if val := reflect.ValueOf(resp); val.Kind() == reflect.Ptr && val.IsNil() {
			resp = reflect.New(val.Type().Elem()).Interface()
			kRes.SetSuccess(resp)
		}
		result := handleServerFailedResp(kRes.GetResult(), errs.ParamError)
		assert.Nil(t, result)
	}

	assert.False(t, ret.Success.Base.Success)
	assert.Equal(t, errs.ParamError.Code(), ret.Success.Base.Code)
	assert.Equal(t, errs.ParamError.Msg(), ret.Success.Base.Msg)
}
