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
		_ = repos.CreateLoginRecord(ctx, &po.LoginRecord{
			AccountID: account.AccountID,
			LoginAt:   time.Now(),
			IPv4:      req.IPv4,
			Device:    req.Device,
			Status:    string(domain.LoginStatusFail),
			Reason:    errs.AccountInvalidError.Msg(),
		})

		return "", errs.AccountInvalidError
	}

	if !utils.PasswordVerify(account.Salt, account.Password, req.Password) {
		// 验证不通过
		_ = repos.CreateLoginRecord(ctx, &po.LoginRecord{
			AccountID: account.AccountID,
			LoginAt:   time.Now(),
			IPv4:      req.IPv4,
			Device:    req.Device,
			Status:    string(domain.LoginStatusFail),
			Reason:    errs.PasswordIncorrect.Msg(),
		})

		return "", errs.PasswordIncorrect
	}

	_ = repos.CreateLoginRecord(ctx, &po.LoginRecord{
		AccountID: account.AccountID,
		LoginAt:   time.Now(),
		IPv4:      req.IPv4,
		Device:    req.Device,
		Status:    string(domain.LoginStatusSuccess),
	})

	return account.AccountID, nil
}

func QueryLoginRecord(ctx context.Context, req *domain.LoginRecordQueryReq) ([]*domain.LoginRecord, int64, errs.Error) {
	limit, offset := req.Size, (req.Page-1)*req.Size
	recordList, total, err := repos.SelectLoginRecord(ctx, req.AccountID, int(limit), int(offset))
	if err != nil {
		return nil, 0, err
	}

	var result []*domain.LoginRecord
	for _, record := range recordList {
		result = append(result, &domain.LoginRecord{
			AccountID: record.AccountID,
			LoginAt:   record.LoginAt,
			IPv4:      record.IPv4,
			Device:    record.Device,
			Reason:    record.Reason,
			Status:    domain.LoginStatus(record.Status),
		})
	}

	return result, total, nil
}
