package rpc

import (
	"context"
	"shield/common/errs"
	"shield/common/middleware/kitex"
	"shield/gateway/biz/model/consts"
	"shield/gateway/kitex_gen/kaidog/shield/account"
	"shield/gateway/kitex_gen/kaidog/shield/account/accountservice"

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
