package rpc

import (
	"context"
	"shield/common/errs"
	"shield/common/middleware/kitex"
	"shield/gateway/biz/model/consts"
	"shield/gateway/biz/model/kaidog/shield/gateway"
	"shield/gateway/kitex_gen/kaidog/shield/account"
	"shield/gateway/kitex_gen/kaidog/shield/account/accountservice"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/kitex/client"
)

var accountClient accountservice.Client

func initAccountClient() {
	accountClient = accountservice.MustNewClient(
		consts.ServiceNameAccount,
		client.WithHostPorts("0.0.0.0:8888"),
		client.WithSuite(kitex.NewClientSuite()),
	)
}

func buildRpcErr(err error) errs.Error {
	if bizErr, ok := err.(errs.Error); ok {
		return bizErr
	}

	return errs.RpcError.SetErr(err)
}

type LoginReq struct {
	Username string
	Password string
	Ipv4     string
	Device   string
}

type LoginResp struct {
	AccountId string
}

type UserInfoQueryReq struct {
	AccountId string
}

type UserInfoQueryResp struct {
	AccountId   string
	UserId      string
	Name        string
	Gender      gateway.Gender
	Phone       string
	Email       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

var genderMapper = map[account.Gender]gateway.Gender{
	account.Gender_male:   gateway.Gender_male,
	account.Gender_female: gateway.Gender_female,
	account.Gender_others: gateway.Gender_others,
}

type LoginRecordQueryResp struct {
	Page       int64
	Size       int64
	Total      int64
	RecordList []*LoginRecord
}

type LoginRecord struct {
	AccountId string
	Ipv4      string
	Device    string
	Status    gateway.LoginStatus
	Reason    string
	LoginAt   time.Time
}

var loginStatusMapper = map[account.LoginStatus]gateway.LoginStatus{
	account.LoginStatus_success: gateway.LoginStatus_success,
	account.LoginStatus_fail:    gateway.LoginStatus_fail,
}

func Login(ctx context.Context, req *LoginReq) (*LoginResp, errs.Error) {
	resp, err := accountClient.Login(ctx,
		&account.AccountLoginReq{
			Username: req.Username,
			Password: req.Password,
			Ipv4:     req.Ipv4,
			Device:   req.Device,
		})
	if err != nil {
		return nil, buildRpcErr(err)
	}

	return &LoginResp{
		AccountId: resp.GetAccountID(),
	}, nil
}

func QueryUserInfoByAccountId(ctx context.Context, accountId string) (*UserInfoQueryResp, errs.Error) {
	resp, err := accountClient.QueryUser(ctx,
		&account.UserQueryReq{
			AccountID: thrift.StringPtr(accountId),
		})
	if err != nil {
		return nil, buildRpcErr(err)
	}

	return &UserInfoQueryResp{
		AccountId:   resp.GetUser().GetAccountID(),
		UserId:      resp.GetUser().GetUserID(),
		Name:        resp.GetUser().GetName(),
		Gender:      genderMapper[resp.GetUser().GetGender()],
		Phone:       resp.GetUser().GetPhone(),
		Email:       resp.GetUser().GetEmail(),
		Description: resp.GetUser().GetDescription(),
		CreatedAt:   time.Unix(resp.GetUser().GetCreatedAt(), 0),
		UpdatedAt:   time.Unix(resp.GetUser().GetUpdatedAt(), 0),
	}, nil
}

func QueryLoginRecordByAccountId(ctx context.Context, accountId string) (*LoginRecordQueryResp, errs.Error) {
	resp, err := accountClient.QueryLoginRecord(ctx,
		&account.LoginRecordQueryReq{
			AccountID: accountId,
			Page:      1,
			Size:      10,
		})
	if err != nil {
		return nil, buildRpcErr(err)
	}

	var recordList []*LoginRecord
	for _, record := range resp.GetRecordList() {
		recordList = append(recordList, &LoginRecord{
			AccountId: record.GetAccountID(),
			Ipv4:      record.GetIpv4(),
			Device:    record.GetDevice(),
			Status:    loginStatusMapper[record.GetStatus()],
			Reason:    record.GetReason(),
			LoginAt:   time.Unix(record.GetLoginAt(), 0),
		})
	}

	return &LoginRecordQueryResp{
		Page:       resp.GetPage(),
		Size:       resp.GetSize(),
		Total:      resp.GetTotal(),
		RecordList: recordList,
	}, nil
}

type UpdatePasswordReq struct {
	AccountId   string
	OldPassword string
	NewPassword string
}

func UpdatePassword(ctx context.Context, req *UpdatePasswordReq) errs.Error {
	_, err := accountClient.UpdateAccountPassword(ctx,
		&account.AccountPasswordUpdateReq{
			AccountID:    req.AccountId,
			Password:     req.OldPassword,
			NewPassword_: req.NewPassword,
		})
	if err != nil {
		return buildRpcErr(err)
	}

	return nil

}
