package mysql

import (
	"context"
	"shield/account/internal/model/po"
	"shield/common/errs"
	"shield/common/logs"
	"shield/common/utils/gorm_utils"

	"gorm.io/gorm"
)

type loginRecordDal struct {
	Dal
}

func NewLoginRecordDal(gormDB ...*gorm.DB) *loginRecordDal {
	dal := &loginRecordDal{
		Dal: NewDefaultDal(gormDB...),
	}

	return dal
}

func (dal *loginRecordDal) Insert(ctx context.Context, loginRecordPO ...*po.LoginRecord) errs.Error {
	if err := dal.GetGormDB().WithContext(ctx).Create(loginRecordPO).Error; err != nil {
		if gorm_utils.IsEntryDuplicateErr(err) {
			return errs.DbDuplicateError.SetErr(err)
		}

		logs.CtxErrorf(ctx, "insert user err: %v", err)
		return errs.DbError.SetErr(err)
	}

	return nil
}

func (dal *loginRecordDal) Select(ctx context.Context, accountID string, limit, offset int) ([]*po.LoginRecord, int64, errs.Error) {
	var result []*po.LoginRecord
	if err := dal.GetGormDB().WithContext(ctx).
		Where("account_id", accountID).
		Order("login_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&result).Error; err != nil {
		logs.CtxErrorf(ctx, "select login record err: %v", err)
		return nil, 0, errs.DbError.SetErr(err)
	}

	var total int64
	if err := dal.GetGormDB().WithContext(ctx).
		Model(po.NewLoginRecord()).
		Where("account_id", accountID).
		Count(&total).Error; err != nil {
		logs.CtxErrorf(ctx, "count login record err: %v", err)
		return nil, 0, errs.DbError.SetErr(err)
	}

	return result, total, nil
}
