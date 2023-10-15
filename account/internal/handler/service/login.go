package service

import (
	"context"
	"shield/account/internal/repos"
	"shield/account/internal/utils"
	"shield/account/model/domain"
	"shield/account/model/po"
	"shield/common/errs"
	"shield/common/logs"
	"time"
)

func Login(ctx context.Context, req *domain.LoginReq) (string, errs.Error) {
	// 1、验证密码有效性
	account, err := repos.SelectAccountByUsername(ctx, req.Username)
	if err != nil {
		return "", err
	}
	if account == nil {
		logs.CtxWarn(ctx, "account not exist")
		return "", errs.AccountNotExistError
	}

	if account.Status != string(domain.AccountStatusValid) {
		// 账户不可用
		err := errs.AccountInvalidError
		defer repos.CreateLoginRecord(ctx, &po.LoginRecord{
			AccountID: account.AccountID,
			LoginAt:   time.Now(),
			IPv4:      req.IPv4,
			Device:    req.Device,
			Status:    string(domain.LoginStatusFail),
			Reason:    err.Msg(),
		})

		return "", err
	}

	if !utils.PasswordVerify(account.Salt, account.Password, req.Password) {
		// 验证不通过
		err := errs.PasswordIncorrect
		defer repos.CreateLoginRecord(ctx, &po.LoginRecord{
			AccountID: account.AccountID,
			LoginAt:   time.Now(),
			IPv4:      req.IPv4,
			Device:    req.Device,
			Status:    string(domain.LoginStatusFail),
			Reason:    err.Msg(),
		})

		return "", err
	}

	defer repos.CreateLoginRecord(ctx, &po.LoginRecord{
		AccountID: account.AccountID,
		LoginAt:   time.Now(),
		IPv4:      req.IPv4,
		Device:    req.Device,
		Status:    string(domain.LoginStatusSuccess),
	})

	return account.AccountID, nil
}

func QueryLoginRecord(ctx context.Context, req *domain.LoginRecordQueryReq)
