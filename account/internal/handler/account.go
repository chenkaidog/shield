package handler

import (
	"context"
	"shield/account/internal/handler/service"
	"shield/account/kitex_gen/kaidog/shield/account"
	"shield/account/model/domain"
	"shield/common/errs"

	"github.com/apache/thrift/lib/go/thrift"
)

func QueryAccount(ctx context.Context, req *account.AccountQueryReq) (*account.AccountQueryResp, errs.Error) {
	resultList, total, err := service.QueryAccount(ctx, &domain.AccountQueryReq{
		Page: req.GetPage(),
		Size: req.GetSize(),
	})
	if err != nil {
		return nil, err
	}

	resp := account.NewAccountQueryResp()
	var accountList []*account.Account
	for _, result := range resultList {
		var status account.AccountStatus
		switch result.Status {
		case domain.AccountStatusValid:
			status = account.AccountStatus_valid
		case domain.AccountStatusInvalid:
			status = account.AccountStatus_invalid
		}
		accountList = append(accountList,
			&account.Account{
				AccountID: result.AccountID,
				Username:  result.Username,
				Status:    status,
			})
	}

	resp.SetAccountList(accountList)
	resp.SetTotal(thrift.Int64Ptr(total))
	resp.SetPage(thrift.Int64Ptr(req.GetPage()))
	resp.SetSize(thrift.Int64Ptr(req.GetSize()))

	return resp, nil
}

func CreateAccount(ctx context.Context, req *account.AccountCreateReq) (*account.AccountCreateResp, errs.Error) {
	acc, err := service.CreateAccount(ctx, &domain.AccountCreateReq{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}

	resp := account.NewAccountCreateResp()
	resp.SetAccountID(&acc.AccountID)
	return resp, nil
}

func UpdateAccountPassword(ctx context.Context, req *account.AccountPasswordUpdateReq) (*account.AccountPasswordUpdateResp, errs.Error) {
	err := service.UpdateAccountPassword(ctx, &domain.AccountPswUpdateReq{
		AccountID:   req.GetAccountID(),
		Password:    req.GetPassword(),
		NewPassword: req.GetNewPassword_(),
	})
	if err != nil {
		return nil, err
	}

	return account.NewAccountPasswordUpdateResp(), nil
}

func ResetAccountPassword(ctx context.Context, req *account.AccountPasswordResetReq) (*account.AccountPasswordResetResp, errs.Error) {
	err := service.ResetAccountPassword(ctx, &domain.AccountPswResetReq{
		AccountID: req.GetAccountID(),
		Password:  req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}

	return account.NewAccountPasswordResetResp(), nil
}

func UpdateAccountStatus(ctx context.Context, req *account.AccountStatusUpdateReq) (*account.AccountStatusUpdateResp, errs.Error) {
	var status domain.AccountStatus
	switch req.GetStatus() {
	case account.AccountStatus_valid:
		status = domain.AccountStatusValid
	case account.AccountStatus_invalid:
		status = domain.AccountStatusInvalid
	}

	err := service.UpdateAccountStatus(ctx, &domain.AccountStatusUpdateReq{
		AccountID: req.GetAccountID(),
		Status:    status,
	})
	if err != nil {
		return nil, err
	}

	return account.NewAccountStatusUpdateResp(), nil
}
