package handler

import (
	"context"
	"shield/account/kitex_gen/kaidog/shield/account"
	"shield/common/errs"
)

func QueryAccount(ctx context.Context, req *account.AccountQueryReq) (*account.AccountQueryResp, errs.Error) {
	return nil, nil
}

func CreateAccount(ctx context.Context, req *account.AccountCreateReq) (*account.AccountCreateResp, errs.Error) {
	return nil, nil
}

func UpdateAccountPassword(ctx context.Context, req *account.AccountPasswordUpdateReq) (*account.AccountPasswordUpdateResp, errs.Error) {
	return nil, nil
}

func ResetAccountPassword(ctx context.Context, req *account.AccountPasswordResetReq) (*account.AccountPasswordResetResp, errs.Error) {
	return nil, nil
}

func UpdateAccountStatus(ctx context.Context, req *account.AccountStatusUpdateReq) (*account.AccountStatusUpdateResp, errs.Error) {
	return nil, nil
}
