package handler

import (
	"context"
	"shield/account/kitex_gen/kaidog/shield/account"
	"shield/common/errs"
)

func Login(ctx context.Context, req *account.AccountLoginReq) (*account.AccountLoginResp, errs.Error) {
	return nil, nil
}

func QueryLoginRecord(ctx context.Context, req *account.LoginRecordQueryReq) (*account.LoginRecordQueryResp, errs.Error) {
	return nil, nil
}
