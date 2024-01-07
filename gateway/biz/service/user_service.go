package service

import (
	"context"
	"shield/common/errs"
	"shield/gateway/biz/model/kaidog/shield/gateway"
	"shield/gateway/biz/repos"
	"shield/gateway/biz/rpc"
)

type LoginReq struct {
	Username string
	Password string
	Ip       string
	Device   string
}

type LoginResp struct {
	AccountID string
}

func Login(ctx context.Context, req *LoginReq) (*LoginResp, errs.Error) {
	rpcResp, bizErr := rpc.Login(
		ctx,
		&rpc.LoginReq{
			Username: req.Username,
			Password: req.Password,
			Ipv4:     req.Ip,
			Device:   req.Device,
		},
	)
	if bizErr != nil {
		return nil, bizErr
	}

	return &LoginResp{
		AccountID: rpcResp.AccountId,
	}, nil
}

func StoreSessionId(ctx context.Context, accountId, sessId string) errs.Error {
	return repos.SetAccountSessionID(ctx, accountId, sessId)
}

func GetSessionId(ctx context.Context, accountId string) (string, errs.Error) {
	return repos.GetAccountSessionID(ctx, accountId)
}

func RemoveSessionId(ctx context.Context, accountId string) errs.Error {
	return repos.RemoveAccountSessionID(ctx, accountId)
}

type SelfUserInfoQueryReq struct {
	AccountID string
}

type SelfUserInfoQueryResp struct {
	Info *gateway.UserInfo
}

func QuerySelfUserInfo(ctx context.Context, req *SelfUserInfoQueryReq) (*SelfUserInfoQueryResp, errs.Error) {
	rpcResp, bizErr := rpc.QueryUserInfoByAccountId(
		ctx,
		req.AccountID,
	)
	if bizErr != nil {
		return nil, bizErr
	}

	return &SelfUserInfoQueryResp{
		Info: &gateway.UserInfo{
			AccountID:   rpcResp.AccountId,
			UserID:      rpcResp.UserId,
			Name:        rpcResp.Name,
			Gender:      rpcResp.Gender,
			Phone:       rpcResp.Phone,
			Email:       rpcResp.Email,
			Description: rpcResp.Description,
			CreatedAt:   rpcResp.CreatedAt.Unix(),
			UpdatedAt:   rpcResp.CreatedAt.Unix(),
		},
	}, nil
}

type SelfLoginRecordQueryReq struct {
	AccountID string
}

type SelfLoginRecordQueryResp struct {
	LoginRecord []*gateway.LoginRecord
	Total       int64
	Page        int64
	Size        int64
}

func QuerySelfLoginRecord(ctx context.Context, req *SelfLoginRecordQueryReq) (*SelfLoginRecordQueryResp, errs.Error) {
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

	return &SelfLoginRecordQueryResp{
		Page:        rpcResp.Page,
		Size:        rpcResp.Size,
		Total:       rpcResp.Total,
		LoginRecord: recordList,
	}, nil
}

type PasswordUpdateReq struct {
	AccountID   string
	OldPassword string
	NewPassword string
}

func UpdatePassword(ctx context.Context, req *PasswordUpdateReq) errs.Error {
	return rpc.UpdatePassword(
		ctx,
		&rpc.UpdatePasswordReq{
			AccountId:   req.AccountID,
			OldPassword: req.OldPassword,
			NewPassword: req.NewPassword,
		},
	)
}
