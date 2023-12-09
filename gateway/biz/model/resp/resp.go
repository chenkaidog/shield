package resp

import (
	"shield/common/errs"

	"shield/gateway/biz/model/kaidog/shield/gateway"
)

type Response struct {
	gateway.BaseResp
	Body interface{} `json:"body,omitempty"`
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

func NewFailResp(code int32, msg string) *Response {
	resp := &Response{}
	resp.Code = code
	resp.Msg = msg
	resp.Success = false

	return resp
}

type UserInfoQueryResp struct {
	AccountId   string         `json:"account_id"`
	UserId      string         `json:"user_id"`
	Name        string         `json:"name"`
	Gender      gateway.Gender `json:"gender"`
	Phone       string         `json:"phone"`
	Email       string         `json:"email"`
	Description string         `json:"description"`
	CreatedAt   int64          `json:"created_at"`
	UpdatedAt   int64          `json:"updated_at"`
}

type LoginRecord struct {
	AccountId string              `json:"account_id"`
	Ipv4      string              `json:"ipv4"`
	Device    string              `json:"device"`
	Status    gateway.LoginStatus `json:"status"`
	Reason    string              `json:"reason"`
	LoginAt   int64               `json:"login_at"`
}

type LoginRecordQueryResp struct {
	Page       int64          `json:"page"`
	Size       int64          `json:"size"`
	Total      int64          `json:"total"`
	RecordList []*LoginRecord `json:"record_list"`
}
