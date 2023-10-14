package main

import (
	"context"
	"shield/account/internal/handler"
	account "shield/account/kitex_gen/kaidog/shield/account"
)

// AccountServiceImpl implements the last service interface defined in the IDL.
type AccountServiceImpl struct{}

// QueryAccount implements the AccountServiceImpl interface.
func (s *AccountServiceImpl) QueryAccount(ctx context.Context, req *account.AccountQueryReq) (resp *account.AccountQueryResp, err error) {
	return handler.QueryAccount(ctx, req)
}

// CreateAccount implements the AccountServiceImpl interface.
func (s *AccountServiceImpl) CreateAccount(ctx context.Context, req *account.AccountCreateReq) (resp *account.AccountCreateResp, err error) {
	return handler.CreateAccount(ctx, req)
}

// UpdateAccountPassword implements the AccountServiceImpl interface.
func (s *AccountServiceImpl) UpdateAccountPassword(ctx context.Context, req *account.AccountPasswordUpdateReq) (resp *account.AccountPasswordUpdateResp, err error) {
	return handler.UpdateAccountPassword(ctx, req)
}

// ResetAccountPassword implements the AccountServiceImpl interface.
func (s *AccountServiceImpl) ResetAccountPassword(ctx context.Context, req *account.AccountPasswordResetReq) (resp *account.AccountPasswordResetResp, err error) {
	return handler.ResetAccountPassword(ctx, req)
}

// UpdateAccountStatus implements the AccountServiceImpl interface.
func (s *AccountServiceImpl) UpdateAccountStatus(ctx context.Context, req *account.AccountStatusUpdateReq) (resp *account.AccountStatusUpdateResp, err error) {
	return handler.UpdateAccountStatus(ctx, req)
}

// Login implements the AccountServiceImpl interface.
func (s *AccountServiceImpl) Login(ctx context.Context, req *account.AccountLoginReq) (resp *account.AccountLoginResp, err error) {
	return handler.Login(ctx, req)
}

// QueryLoginRecord implements the AccountServiceImpl interface.
func (s *AccountServiceImpl) QueryLoginRecord(ctx context.Context, req *account.LoginRecordQueryReq) (resp *account.LoginRecordQueryResp, err error) {
	return handler.QueryLoginRecord(ctx, req)
}

// CreateUser implements the AccountServiceImpl interface.
func (s *AccountServiceImpl) CreateUser(ctx context.Context, req *account.UserCreateReq) (resp *account.UserCreateResp, err error) {
	return handler.CreateUser(ctx, req)
}

// QueryUser implements the AccountServiceImpl interface.
func (s *AccountServiceImpl) QueryUser(ctx context.Context, req *account.UserQueryReq) (resp *account.UserQueryResp, err error) {
	return handler.QueryUser(ctx, req)
}

// UpdateUser implements the AccountServiceImpl interface.
func (s *AccountServiceImpl) UpdateUser(ctx context.Context, req *account.UserUpdateReq) (resp *account.UserUpdateResp, err error) {
	return handler.UpdateUser(ctx, req)
}
