// Code generated by hertz generator.

package gateway

import (
	"context"

	"shield/common/errs"
	"shield/common/logs"
	gateway "shield/gateway/biz/model/kaidog/shield/gateway"
	"shield/gateway/biz/service"
	"shield/gateway/biz/util"

	"github.com/cloudwego/hertz/pkg/app"
)

// CreateAccount .
// @router /create_account [POST]
func CreateAccount(ctx context.Context, c *app.RequestContext) {
	var err error
	var req gateway.AccountCreateReq
	var resp gateway.AccountCreateResp
	err = c.BindAndValidate(&req)
	if err != nil {
		logs.CtxErrorf(ctx, "BindAndValidate fail, %v", err)
		util.BuildBizResp(c, &resp, errs.ParamError.SetErr(err))
		return
	}

	createResp, bizErr := service.CreateAccount(
		ctx,
		&service.AccountCreateReq{
			Username: req.GetUsername(),
			Password: req.GetPassword(),
		},
	)
	if createResp != nil {
		resp.AccountID = createResp.AccountID
	}

	util.BuildBizResp(c, &resp, bizErr)
}

// CreateUser .
// @router /create_user [POST]
func CreateUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req gateway.UserCreateReq
	var resp gateway.UserCreateResp
	err = c.BindAndValidate(&req)
	if err != nil {
		logs.CtxErrorf(ctx, "BindAndValidate fail, %v", err)
		util.BuildBizResp(c, &resp, errs.ParamError.SetErr(err))
		return
	}

	createResp, bizErr := service.CreateUser(
		ctx,
		&service.UserCreateReq{
			AccountID:   req.GetAccountID(),
			Name:        req.GetName(),
			Gender:      req.GetGender(),
			Phone:       req.GetPhone(),
			Email:       req.GetEmail(),
			Description: req.GetDescription(),
		},
	)
	if createResp != nil {
		resp.UserID = createResp.UserID
	}

	util.BuildBizResp(c, &resp, bizErr)
}

// UpdateUserInfo .
// @router /update_user [POST]
func UpdateUserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req gateway.UserInfoUpdateReq
	var resp gateway.UserInfoUpdateResp
	err = c.BindAndValidate(&req)
	if err != nil {
		logs.CtxErrorf(ctx, "BindAndValidate fail, %v", err)
		util.BuildBizResp(c, &resp, errs.ParamError.SetErr(err))
		return
	}

	bizErr := service.UpdateUserInfo(
		ctx,
		&service.UserInfoUpdateReq{
			UserID:      req.GetUserID(),
			Name:        req.GetName(),
			Gender:      req.GetGender(),
			Phone:       req.GetPhone(),
			Email:       req.GetPhone(),
			Description: req.GetDescription(),
		},
	)

	util.BuildBizResp(c, &resp, bizErr)
}

// ResetPassword .
// @router /rest_password [POST]
func ResetPassword(ctx context.Context, c *app.RequestContext) {
	var err error
	var req gateway.PasswordRestReq
	var resp gateway.PasswordRestResp
	err = c.BindAndValidate(&req)
	if err != nil {
		logs.CtxErrorf(ctx, "BindAndValidate fail, %v", err)
		util.BuildBizResp(c, &resp, errs.ParamError.SetErr(err))
		return
	}

	bizErr := service.ResetPassword(
		ctx,
		&service.PasswordResetReq{
			AccountID: req.GetAccountID(),
			Password:  req.GetNewPassword(),
		},
	)

	util.BuildBizResp(c, &resp, bizErr)
}

// SwitchAccountStatus .
// @router /switch_account_status [POST]
func SwitchAccountStatus(ctx context.Context, c *app.RequestContext) {
	var err error
	var req gateway.AccountStatusSwitchReq
	var resp gateway.AccountStatusSwitchResp
	err = c.BindAndValidate(&req)
	if err != nil {
		logs.CtxErrorf(ctx, "BindAndValidate fail, %v", err)
		util.BuildBizResp(c, &resp, errs.ParamError.SetErr(err))
		return
	}

	bizErr := service.ChangeAccountStatus(
		ctx,
		&service.AccountStatusChangeReq{
			AccountID: req.GetAccountID(),
			Status:    req.GetStatus(),
		},
	)

	util.BuildBizResp(c, &resp, bizErr)
}

// QueryAccount .
// @router /admin/query_account [GET]
func QueryAccount(ctx context.Context, c *app.RequestContext) {
	var err error
	var req gateway.AccountQueryReq
	var resp gateway.AccountQueryResp
	err = c.BindAndValidate(&req)
	if err != nil {
		logs.CtxErrorf(ctx, "BindAndValidate fail, %v", err)
		util.BuildBizResp(c, &resp, errs.ParamError.SetErr(err))
		return
	}

	queryResp, bizErr := service.QueryAccount(
		ctx,
		&service.AccountQueryReq{
			Page: req.GetPage(),
			Size: req.GetSize(),
		},
	)
	if queryResp != nil {
		resp.AccountList = queryResp.AccountList
		resp.Page = queryResp.Page
		resp.Size = queryResp.Size
		resp.Total = queryResp.Total
	}

	util.BuildBizResp(c, &resp, bizErr)
}

// QueryUserInfo .
// @router /operator/admin/query_user_info [GET]
func QueryUserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req gateway.UserInfoQueryReq
	var resp gateway.UserInfoQueryResp
	err = c.BindAndValidate(&req)
	if err != nil {
		logs.CtxErrorf(ctx, "BindAndValidate fail, %v", err)
		util.BuildBizResp(c, &resp, errs.ParamError.SetErr(err))
		return
	}

	queryResp, bizErr := service.QueryUserInfo(
		ctx,
		&service.UserInfoQueryReq{
			AccountIdList: req.AccountIdList,
		},
	)
	if queryResp != nil {
		resp.UserList = queryResp.UserList
	}

	util.BuildBizResp(c, &resp, bizErr)
}

// QueryLoginRecord .
// @router /operator/admin/query_login_record [GET]
func QueryLoginRecord(ctx context.Context, c *app.RequestContext) {
	var err error
	var req gateway.LoginRecordQueryReq
	var resp gateway.LoginRecordQueryResp
	err = c.BindAndValidate(&req)
	if err != nil {
		logs.CtxErrorf(ctx, "BindAndValidate fail, %v", err)
		util.BuildBizResp(c, &resp, errs.ParamError.SetErr(err))
		return
	}

	queryResp, bizErr := service.QueryLoginRecord(
		ctx,
		&service.LoginRecordQueryReq{
			AccountID: req.GetAccountID(),
		},
	)
	if queryResp != nil {
		resp.LoginRecord = queryResp.LoginRecord
		resp.Total = queryResp.Total
		resp.Page = queryResp.Page
		resp.Size = queryResp.Size
	}

	util.BuildBizResp(c, &resp, bizErr)
}
