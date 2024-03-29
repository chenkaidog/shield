// Code generated by Kitex v0.7.3. DO NOT EDIT.

package accountservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	account "shield/gateway/kitex_gen/kaidog/shield/account"
)

func serviceInfo() *kitex.ServiceInfo {
	return accountServiceServiceInfo
}

var accountServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "AccountService"
	handlerType := (*account.AccountService)(nil)
	methods := map[string]kitex.MethodInfo{
		"QueryAccount":          kitex.NewMethodInfo(queryAccountHandler, newAccountServiceQueryAccountArgs, newAccountServiceQueryAccountResult, false),
		"CreateAccount":         kitex.NewMethodInfo(createAccountHandler, newAccountServiceCreateAccountArgs, newAccountServiceCreateAccountResult, false),
		"UpdateAccountPassword": kitex.NewMethodInfo(updateAccountPasswordHandler, newAccountServiceUpdateAccountPasswordArgs, newAccountServiceUpdateAccountPasswordResult, false),
		"ResetAccountPassword":  kitex.NewMethodInfo(resetAccountPasswordHandler, newAccountServiceResetAccountPasswordArgs, newAccountServiceResetAccountPasswordResult, false),
		"UpdateAccountStatus":   kitex.NewMethodInfo(updateAccountStatusHandler, newAccountServiceUpdateAccountStatusArgs, newAccountServiceUpdateAccountStatusResult, false),
		"Login":                 kitex.NewMethodInfo(loginHandler, newAccountServiceLoginArgs, newAccountServiceLoginResult, false),
		"QueryLoginRecord":      kitex.NewMethodInfo(queryLoginRecordHandler, newAccountServiceQueryLoginRecordArgs, newAccountServiceQueryLoginRecordResult, false),
		"CreateUser":            kitex.NewMethodInfo(createUserHandler, newAccountServiceCreateUserArgs, newAccountServiceCreateUserResult, false),
		"QueryUser":             kitex.NewMethodInfo(queryUserHandler, newAccountServiceQueryUserArgs, newAccountServiceQueryUserResult, false),
		"UpdateUser":            kitex.NewMethodInfo(updateUserHandler, newAccountServiceUpdateUserArgs, newAccountServiceUpdateUserResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "account",
		"ServiceFilePath": `../common/idl/account.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.7.3",
		Extra:           extra,
	}
	return svcInfo
}

func queryAccountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*account.AccountServiceQueryAccountArgs)
	realResult := result.(*account.AccountServiceQueryAccountResult)
	success, err := handler.(account.AccountService).QueryAccount(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAccountServiceQueryAccountArgs() interface{} {
	return account.NewAccountServiceQueryAccountArgs()
}

func newAccountServiceQueryAccountResult() interface{} {
	return account.NewAccountServiceQueryAccountResult()
}

func createAccountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*account.AccountServiceCreateAccountArgs)
	realResult := result.(*account.AccountServiceCreateAccountResult)
	success, err := handler.(account.AccountService).CreateAccount(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAccountServiceCreateAccountArgs() interface{} {
	return account.NewAccountServiceCreateAccountArgs()
}

func newAccountServiceCreateAccountResult() interface{} {
	return account.NewAccountServiceCreateAccountResult()
}

func updateAccountPasswordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*account.AccountServiceUpdateAccountPasswordArgs)
	realResult := result.(*account.AccountServiceUpdateAccountPasswordResult)
	success, err := handler.(account.AccountService).UpdateAccountPassword(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAccountServiceUpdateAccountPasswordArgs() interface{} {
	return account.NewAccountServiceUpdateAccountPasswordArgs()
}

func newAccountServiceUpdateAccountPasswordResult() interface{} {
	return account.NewAccountServiceUpdateAccountPasswordResult()
}

func resetAccountPasswordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*account.AccountServiceResetAccountPasswordArgs)
	realResult := result.(*account.AccountServiceResetAccountPasswordResult)
	success, err := handler.(account.AccountService).ResetAccountPassword(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAccountServiceResetAccountPasswordArgs() interface{} {
	return account.NewAccountServiceResetAccountPasswordArgs()
}

func newAccountServiceResetAccountPasswordResult() interface{} {
	return account.NewAccountServiceResetAccountPasswordResult()
}

func updateAccountStatusHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*account.AccountServiceUpdateAccountStatusArgs)
	realResult := result.(*account.AccountServiceUpdateAccountStatusResult)
	success, err := handler.(account.AccountService).UpdateAccountStatus(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAccountServiceUpdateAccountStatusArgs() interface{} {
	return account.NewAccountServiceUpdateAccountStatusArgs()
}

func newAccountServiceUpdateAccountStatusResult() interface{} {
	return account.NewAccountServiceUpdateAccountStatusResult()
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*account.AccountServiceLoginArgs)
	realResult := result.(*account.AccountServiceLoginResult)
	success, err := handler.(account.AccountService).Login(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAccountServiceLoginArgs() interface{} {
	return account.NewAccountServiceLoginArgs()
}

func newAccountServiceLoginResult() interface{} {
	return account.NewAccountServiceLoginResult()
}

func queryLoginRecordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*account.AccountServiceQueryLoginRecordArgs)
	realResult := result.(*account.AccountServiceQueryLoginRecordResult)
	success, err := handler.(account.AccountService).QueryLoginRecord(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAccountServiceQueryLoginRecordArgs() interface{} {
	return account.NewAccountServiceQueryLoginRecordArgs()
}

func newAccountServiceQueryLoginRecordResult() interface{} {
	return account.NewAccountServiceQueryLoginRecordResult()
}

func createUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*account.AccountServiceCreateUserArgs)
	realResult := result.(*account.AccountServiceCreateUserResult)
	success, err := handler.(account.AccountService).CreateUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAccountServiceCreateUserArgs() interface{} {
	return account.NewAccountServiceCreateUserArgs()
}

func newAccountServiceCreateUserResult() interface{} {
	return account.NewAccountServiceCreateUserResult()
}

func queryUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*account.AccountServiceQueryUserArgs)
	realResult := result.(*account.AccountServiceQueryUserResult)
	success, err := handler.(account.AccountService).QueryUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAccountServiceQueryUserArgs() interface{} {
	return account.NewAccountServiceQueryUserArgs()
}

func newAccountServiceQueryUserResult() interface{} {
	return account.NewAccountServiceQueryUserResult()
}

func updateUserHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*account.AccountServiceUpdateUserArgs)
	realResult := result.(*account.AccountServiceUpdateUserResult)
	success, err := handler.(account.AccountService).UpdateUser(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAccountServiceUpdateUserArgs() interface{} {
	return account.NewAccountServiceUpdateUserArgs()
}

func newAccountServiceUpdateUserResult() interface{} {
	return account.NewAccountServiceUpdateUserResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) QueryAccount(ctx context.Context, req *account.AccountQueryReq) (r *account.AccountQueryResp, err error) {
	var _args account.AccountServiceQueryAccountArgs
	_args.Req = req
	var _result account.AccountServiceQueryAccountResult
	if err = p.c.Call(ctx, "QueryAccount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateAccount(ctx context.Context, req *account.AccountCreateReq) (r *account.AccountCreateResp, err error) {
	var _args account.AccountServiceCreateAccountArgs
	_args.Req = req
	var _result account.AccountServiceCreateAccountResult
	if err = p.c.Call(ctx, "CreateAccount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateAccountPassword(ctx context.Context, req *account.AccountPasswordUpdateReq) (r *account.AccountPasswordUpdateResp, err error) {
	var _args account.AccountServiceUpdateAccountPasswordArgs
	_args.Req = req
	var _result account.AccountServiceUpdateAccountPasswordResult
	if err = p.c.Call(ctx, "UpdateAccountPassword", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ResetAccountPassword(ctx context.Context, req *account.AccountPasswordResetReq) (r *account.AccountPasswordResetResp, err error) {
	var _args account.AccountServiceResetAccountPasswordArgs
	_args.Req = req
	var _result account.AccountServiceResetAccountPasswordResult
	if err = p.c.Call(ctx, "ResetAccountPassword", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateAccountStatus(ctx context.Context, req *account.AccountStatusUpdateReq) (r *account.AccountStatusUpdateResp, err error) {
	var _args account.AccountServiceUpdateAccountStatusArgs
	_args.Req = req
	var _result account.AccountServiceUpdateAccountStatusResult
	if err = p.c.Call(ctx, "UpdateAccountStatus", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, req *account.AccountLoginReq) (r *account.AccountLoginResp, err error) {
	var _args account.AccountServiceLoginArgs
	_args.Req = req
	var _result account.AccountServiceLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryLoginRecord(ctx context.Context, req *account.LoginRecordQueryReq) (r *account.LoginRecordQueryResp, err error) {
	var _args account.AccountServiceQueryLoginRecordArgs
	_args.Req = req
	var _result account.AccountServiceQueryLoginRecordResult
	if err = p.c.Call(ctx, "QueryLoginRecord", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CreateUser(ctx context.Context, req *account.UserCreateReq) (r *account.UserCreateResp, err error) {
	var _args account.AccountServiceCreateUserArgs
	_args.Req = req
	var _result account.AccountServiceCreateUserResult
	if err = p.c.Call(ctx, "CreateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryUser(ctx context.Context, req *account.UserQueryReq) (r *account.UserQueryResp, err error) {
	var _args account.AccountServiceQueryUserArgs
	_args.Req = req
	var _result account.AccountServiceQueryUserResult
	if err = p.c.Call(ctx, "QueryUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateUser(ctx context.Context, req *account.UserUpdateReq) (r *account.UserUpdateResp, err error) {
	var _args account.AccountServiceUpdateUserArgs
	_args.Req = req
	var _result account.AccountServiceUpdateUserResult
	if err = p.c.Call(ctx, "UpdateUser", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
