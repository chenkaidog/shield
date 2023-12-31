package mysql

import (
	"context"
	"shield/account/internal/model/po"
	"shield/common/errs"
	"shield/common/logs"
	"shield/common/utils/gorm_utils"

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
		if gorm_utils.IsEntryDuplicateErr(err) {
			return errs.DbDuplicateError.SetErr(err)
		}

		logs.CtxErrorf(ctx, "insert account err: %v", err)
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
		logs.CtxErrorf(ctx, "update account err: %v", err)
		return errs.DbError.SetErr(err)
	}

	return nil
}

func (dal *accountDal) Select(ctx context.Context, limit, offset int) ([]*po.Account, int64, errs.Error) {
	var result []*po.Account
	if err := dal.GetGormDB().WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&result).Error; err != nil {
		logs.CtxErrorf(ctx, "select account by id err: %v", err)
		return nil, 0, errs.DbError.SetErr(err)
	}

	var total int64
	if err := dal.GetGormDB().WithContext(ctx).
		Model(po.NewAccount()).
		Count(&total).Error; err != nil {
		logs.CtxErrorf(ctx, "count account err: %v", err)
		return nil, 0, errs.DbError.SetErr(err)
	}

	return result, total, nil
}

func (dal *accountDal) SelectByID(ctx context.Context, accontID string) (*po.Account, errs.Error) {
	result := po.NewAccount()
	if err := dal.GetGormDB().WithContext(ctx).
		Where("account_id", accontID).
		Take(result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logs.CtxErrorf(ctx, "select account by username err: %v", err)
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
		logs.CtxErrorf(ctx, "select account by username err: %v", err)
		return nil, errs.DbError.SetErr(err)
	}

	return result, nil
}
