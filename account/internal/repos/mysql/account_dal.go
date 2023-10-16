package mysql

import (
	"context"
	"shield/account/model/po"
	"shield/common/errs"
	"shield/common/logs"
	"shield/common/utils"

	"gorm.io/gorm"
)

type accountDal struct{ Dal }

func NewAccountDal(gormDB ...*gorm.DB) *accountDal {
	dal := &accountDal{
		Dal: NewDefaultDal(gormDB...),
	}

	return dal
}

func (dal *accountDal) Insert(ctx context.Context, accountPO ...*po.Account) errs.Error {
	if err := dal.GetGormDB().WithContext(ctx).Create(accountPO).Error; err != nil {
		if utils.IsEntryDuplicateErr(err) {
			return errs.DbDuplicateError.SetErr(err)
		}

		logs.CtxError(ctx, "insert account err: %v", err)
		return errs.DbError.SetErr(err)
	}

	return nil
}

func (dal *accountDal) Update(ctx context.Context, accountPO *po.Account) errs.Error {
	err := dal.GetGormDB().WithContext(ctx).
		Omit("account_id").
		Where("account_id", accountPO.AccountID).
		Updates(accountPO).Error
	if err != nil {
		logs.CtxError(ctx, "update account err: %v", err)
		return errs.DbError.SetErr(err)
	}

	return nil
}

func (dal *accountDal) SelectByID(ctx context.Context, accountID string) (*po.Account, errs.Error) {
	result := po.NewAccount()
	if err := dal.GetGormDB().WithContext(ctx).
		Where("account_id", accountID).
		Take(result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logs.CtxError(ctx, "select account by id err: %v", err)
		return nil, errs.DbError.SetErr(err)
	}

	return result, nil
}

func (dal *accountDal) SelectByUsername(ctx context.Context, username string) (*po.Account, errs.Error) {
	result := po.NewAccount()
	if err := dal.GetGormDB().WithContext(ctx).
		Where("username", username).
		Take(result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logs.CtxError(ctx, "select account by username err: %v", err)
		return nil, errs.DbError.SetErr(err)
	}

	return result, nil
}
