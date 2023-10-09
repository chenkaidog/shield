package main

import (
	"context"
	account "shield/account/kitex_gen/kaidog/shield/account"
)

// AccountServiceImpl implements the last service interface defined in the IDL.
type AccountServiceImpl struct{}

// QueryAccount implements the AccountServiceImpl interface.
func (s *AccountServiceImpl) QueryAccount(ctx context.Context, req *account.AccountQueryReq) (resp *account.AccountQueryResp, err error) {
	// TODO: Your code here...
	return
}

// CreateAccount implements the AccountServiceImpl interface.
func (s *AccountServiceImpl) CreateAccount(ctx context.Context, req *account.AccountCreateReq) (resp *account.AccountCreateResp, err error) {
	// TODO: Your code here...
	return
}

// UpdateAccountPassword implements the AccountServiceImpl interface.
func (s *AccountServiceImpl) UpdateAccountPassword(ctx context.Context, req *account.AccountPasswordUpdateReq) (resp *account.AccountPasswordUpdateResp, err error) {
	// TODO: Your code here...
	return
}

// ResetAccountPassword implements the AccountServiceImpl interface.
func (s *AccountServiceImpl) ResetAccountPassword(ctx context.Context, req *account.AccountPasswordResetReq) (resp *account.AccountPasswordResetResp, err error) {
	// TODO: Your code here...
	return
}

// UpdateAccountStatus implements the AccountServiceImpl interface.
func (s *AccountServiceImpl) UpdateAccountStatus(ctx context.Context, req *account.AccountStatusUpdateReq) (resp *account.AccountStatusUpdateResp, err error) {
	// TODO: Your code here...
	return
}

// Login implements the AccountServiceImpl interface.
func (s *AccountServiceImpl) Login(ctx context.Context, req *account.AccountLoginReq) (resp *account.AccountLoginResp, err error) {
	panic("test panic")
	return
}

// QueryLoginRecord implements the AccountServiceImpl interface.
func (s *AccountServiceImpl) QueryLoginRecord(ctx context.Context, req *account.LoginRecordQueryReq) (resp *account.LoginRecordQueryResp, err error) {
	// TODO: Your code here...
	return
}

// CreateUser implements the AccountServiceImpl interface.
func (s *AccountServiceImpl) CreateUser(ctx context.Context, req *account.UserCreateReq) (resp *account.UserCreateResp, err error) {
	// TODO: Your code here...
	return
}

// QueryUser implements the AccountServiceImpl interface.
func (s *AccountServiceImpl) QueryUser(ctx context.Context, req *account.UserQueryReq) (resp *account.UserQueryResp, err error) {
	// TODO: Your code here...
	return
}

// UpdateUser implements the AccountServiceImpl interface.
func (s *AccountServiceImpl) UpdateUser(ctx context.Context, req *account.UserUpdateReq) (resp *account.UserUpdateResp, err error) {
	// TODO: Your code here...
	return
}
