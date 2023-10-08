package errs

import "fmt"

type Error interface {
	error
	Code() int32
	Msg() string
	SetErr(err error) Error
	SetMsg(msg string) Error
}

type BizError struct {
	code  int32
	msg   string
}

func (err *BizError) Error() string {
	return fmt.Sprintf("biz_rror[%d:%s]", err.code, err.msg)
}

func (err *BizError) Code() int32 {
	return err.code
}

func (err *BizError) Msg() string {
	return err.msg
}

func (bizErr *BizError) SetErr(err error) Error {
	return New(bizErr.Code(), err.Error())
}

func (bizErr *BizError) SetMsg(msg string) Error {
	return New(bizErr.Code(),msg)
}

func New(code int32, msg string) Error {
	return &BizError{
		code: code,
		msg: msg,
	}
}

func ErrorEqual(err1, err2 Error) bool {
	// 都为空
	if err1 == nil && err2 == nil {
		return true
	}

	// 只有一个不为空
	if err1 == nil || err2 == nil {
		return false
	}

	// 都不为空
	return err1.Code() == err2.Code()
}

var (
	Success = New(0, "success")
	ServerError = New(1_000_01, "internal server error")
	ParamError  = New(1_000_02, "param error")
)