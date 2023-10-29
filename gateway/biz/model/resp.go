package model

import (
	"shield/common/errs"

	"shield/gateway/biz/model/kaidog/shield/gateway"
)

type Response struct {
	gateway.BaseResp
	Body interface{} `json:"body"`
}

func NewSuccessResp(body interface{}) *Response {
	resp := &Response{
		Body: body,
	}

	resp.Code = errs.Success.Code()
	resp.Msg = errs.Success.Msg()
	resp.Success = true

	return resp
}

func NewFailResp(code int32, msg string)  *Response {
	resp := &Response{} 
	resp.Code = code
	resp.Msg = msg
	resp.Success = false

	return resp
}
