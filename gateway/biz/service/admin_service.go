package service

import (
	"context"
	"shield/common/errs"
	"shield/gateway/biz/model/kaidog/shield/gateway"
	"shield/gateway/biz/rpc"
)

type AccountCreateReq struct {
	Username string
	Password string
}

type AccountCreateResp struct {
	AccountID string
}

func CreateAccount(ctx context.Context, req *AccountCreateReq) (*AccountCreateResp, errs.Error) {
	rpcResp, bizErr := rpc.CreateAccount(
		ctx,
		&rpc.AccountCreateReq{
			Username: req.Username,
			Password: req.Password,
		},
	)
	if bizErr != nil {
		return nil, bizErr
	}

	return &AccountCreateResp{
		AccountID: rpcResp.AccountId,
	}, nil
}

type UserCreateReq struct {
	AccountID   string
	Name        string
	Gender      gateway.Gender
	Phone       string
	Email       string
	Description string
}

type UserCreateResp struct {
	UserID string
}

func CreateUser(ctx context.Context, req *UserCreateReq) (*UserCreateResp, errs.Error) {
	rpcResp, bizErr := rpc.CreateUser(
		ctx,
		&rpc.UserCreateReq{
			AccountId:   req.AccountID,
			Name:        req.Name,
			Gender:      req.Gender,
			Phone:       req.Phone,
			Email:       req.Email,
			Description: req.Description,
		},
	)
	if bizErr != nil {
		return nil, bizErr
	}

	return &UserCreateResp{
		UserID: rpcResp.UserId,
	}, nil
}

type UserInfoUpdateReq struct {
	UserID      string
	Name        string
	Gender      gateway.Gender
	Phone       string
	Email       string
	Description string
}

func UpdateUserInfo(ctx context.Context, req *UserInfoUpdateReq) errs.Error {
	return rpc.UpdateUserInfo(
		ctx,
		&rpc.UserInfoUpdateReq{
			UserId:      req.UserID,
			Name:        req.Name,
			Gender:      req.Gender,
			Phone:       req.Phone,
			Email:       req.Email,
			Description: req.Description,
		},
	)
}

type PasswordResetReq struct {
	AccountID string
	Password  string
}

func ResetPassword(ctx context.Context, req *PasswordResetReq) errs.Error {
	return rpc.ResetPassword(
		ctx,
		&rpc.PasswordResetReq{
			AccountId: req.AccountID,
			Password:  req.Password,
		},
	)
}

type AccountStatusChangeReq struct {
	AccountID string
	Status    gateway.AccountStatus
}

func ChangeAccountStatus(ctx context.Context, req *AccountStatusChangeReq) errs.Error {
	return rpc.SwitchAccountStatus(
		ctx,
		&rpc.AccountStatusSwitchReq{
			AccountId: req.AccountID,
			Status:    req.Status,
		},
	)
}

type AccountQueryReq struct {
	Page int64
	Size int64
}

type AccountQueryResp struct {
	AccountList []*gateway.Account
	Total       int64
	Page        int64
	Size        int64
}

func QueryAccount(ctx context.Context, req *AccountQueryReq) (*AccountQueryResp, errs.Error) {
	rpcResp, bizErr := rpc.QueryAccount(
		ctx,
		&rpc.AccountQueryReq{
			Page: req.Page,
			Size: req.Size,
		},
	)
	if bizErr != nil {
		return nil, bizErr
	}

	var accountList []*gateway.Account
	for _, accountResp := range rpcResp.AccountList {
		accountList = append(
			accountList,
			&gateway.Account{
				AccountID: accountResp.AccountId,
				Username:  accountResp.Username,
				Status:    accountResp.Status,
			})
	}

	return &AccountQueryResp{
		Total:       rpcResp.Total,
		AccountList: accountList,
		Page:        req.Page,
		Size:        req.Size,
	}, nil
}

type UserInfoQueryReq struct {
	AccountIdList []string
}

type UserInfoQueryResp struct {
	UserList []*gateway.UserInfo
}

func QueryUserInfo(ctx context.Context, req *UserInfoQueryReq) (*UserInfoQueryResp, errs.Error) {
	rpcResp, bizErr := rpc.QueryUserInfoByAccountId(
		ctx,
		req.AccountIdList[0],
	)
	if bizErr != nil {
		return nil, bizErr
	}

	userInfo := &gateway.UserInfo{
		AccountID:   rpcResp.AccountId,
		UserID:      rpcResp.UserId,
		Name:        rpcResp.Name,
		Gender:      rpcResp.Gender,
		Phone:       rpcResp.Phone,
		Email:       rpcResp.Email,
		Description: rpcResp.Description,
		CreatedAt:   rpcResp.CreatedAt.Unix(),
		UpdatedAt:   rpcResp.CreatedAt.Unix(),
	}

	return &UserInfoQueryResp{
		UserList: []*gateway.UserInfo{
			userInfo,
		},
	}, nil
}

type LoginRecordQueryReq struct {
	AccountID string
}

type LoginRecordQueryResp struct {
	LoginRecord []*gateway.LoginRecord
	Total       int64
	Page        int64
	Size        int64
}

func QueryLoginRecord(ctx context.Context, req *LoginRecordQueryReq) (*LoginRecordQueryResp, errs.Error) {
	rpcResp, bizErr := rpc.QueryLoginRecordByAccountId(ctx, req.AccountID)
	if bizErr != nil {
		return nil, bizErr
	}

	var recordList []*gateway.LoginRecord
	for _, record := range rpcResp.RecordList {
		recordList = append(
			recordList,
			&gateway.LoginRecord{
				AccountID: record.AccountId,
				Ipv4:      record.Ipv4,
				Device:    record.Device,
				Status:    record.Status,
				Reason:    record.Reason,
				LoginAt:   record.LoginAt.Unix(),
			},
		)
	}

	return &LoginRecordQueryResp{
		Page:        rpcResp.Page,
		Size:        rpcResp.Size,
		Total:       rpcResp.Total,
		LoginRecord: recordList,
	}, nil
}
