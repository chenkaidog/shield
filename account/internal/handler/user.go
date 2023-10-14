package handler

import (
	"context"
	"shield/account/kitex_gen/kaidog/shield/account"
	"shield/common/errs"
)

func CreateUser(ctx context.Context, req *account.UserCreateReq) (resp *account.UserCreateResp, err errs.Error) {
	return nil, nil
}

func  QueryUser(ctx context.Context, req *account.UserQueryReq) (resp *account.UserQueryResp, err errs.Error) {
	return nil, nil
}

func  UpdateUser(ctx context.Context, req *account.UserUpdateReq) (resp *account.UserUpdateResp, err errs.Error) {
	return nil, nil
}