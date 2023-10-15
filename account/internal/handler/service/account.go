package service

import (
	"context"
	"shield/account/internal/repos"
	"shield/account/internal/utils"
	"shield/account/model/domain"
	"shield/account/model/po"
	"shield/common/errs"
	"shield/common/logs"
	common_utils "shield/common/utils"
)

func CreateAccount(ctx context.Context, req *domain.AccountCreateReq) (*domain.Account, errs.Error) {
	// 1、判断用户名是否已存在
	account, err := repos.SelectAccountByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	if account != nil {
		logs.CtxWarn(ctx, "username already exists")
		return nil, errs.UsernameDuplidateError
	}

	// 2、密码编码并创建账号
	salt, password := utils.EncodePassword(req.Password)
	accountID := common_utils.NewUUID()
	defaultStatus := domain.AccountStatusValid
	if err = repos.CreateAccount(ctx, &po.Account{
		AccountID: accountID,
		Username:  req.Username,
		Salt:      salt,
		Password:  password,
		Status:    string(defaultStatus),
	}); err != nil {
		if errs.ErrorEqual(err, errs.DbDuplicateError) {
			return nil, errs.UsernameDuplidateError
		}
		return nil, err
	}

	return &domain.Account{
		AccountID: accountID,
		Username:  req.Username,
		Status:    defaultStatus,
	}, nil
}

func UpdateAccountPassword(ctx context.Context, req *domain.AccountPswUpdateReq) errs.Error {
	// 1、验证密码有效性
	account, err := repos.SelectAccountByID(ctx, req.AccountID)
	if err != nil {
		return err
	}
	if account == nil {
		logs.CtxWarn(ctx, "account not exist")
		return errs.AccountNotExistError
	}

	if !utils.PasswordVerify(account.Salt, account.Password, req.Password) {
		// 验证不通过
		logs.CtxInfo(ctx, "password incorrect")
		return errs.PasswordIncorrect
	}

	// 2、修改密码
	salt, password := utils.EncodePassword(req.Password)
	return repos.UpdateAccount(ctx, &po.Account{
		AccountID: req.AccountID,
		Salt:      salt,
		Password:  password,
	})
}

func ResetAccountPassword(ctx context.Context, req *domain.AccountPswResetReq) errs.Error {
	salt, password := utils.EncodePassword(req.Password)
	return repos.UpdateAccount(ctx, &po.Account{
		AccountID: req.AccountID,
		Salt:      salt,
		Password:  password,
	})
}

func UpdateAccountStatus(ctx context.Context, req *domain.AccountStatusUpdateReq) errs.Error {
	return repos.UpdateAccount(ctx, &po.Account{
		AccountID: req.AccountID,
		Status:    string(req.Status),
	})
}

func QueryAccount(ctx context.Context, req *domain.AccountQueryReq) (*domain.Account, errs.Error) {
	account, err := repos.SelectAccountByID(ctx, req.AccountID)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return nil, nil
	}

	return &domain.Account{
		AccountID: account.AccountID,
		Username:  account.Username,
		Status:    domain.AccountStatus(account.Status),
	}, nil
}
