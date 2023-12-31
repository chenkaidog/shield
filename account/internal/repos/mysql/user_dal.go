package mysql

import (
	"context"
	"shield/account/internal/model/po"
	"shield/common/errs"
	"shield/common/logs"
	"shield/common/utils/gorm_utils"

	"gorm.io/gorm"
)

type userDal struct {
	Dal
}

func NewUserDal(gormDB ...*gorm.DB) *userDal {
	dal := &userDal{
		Dal: NewDefaultDal(gormDB...),
	}

	return dal
}

func (dal *userDal) Insert(ctx context.Context, userPO ...*po.User) errs.Error {
	if err := dal.GetGormDB().WithContext(ctx).Create(userPO).Error; err != nil {
		if gorm_utils.IsEntryDuplicateErr(err) {
			return errs.DbDuplicateError.SetErr(err)
		}

		logs.CtxErrorf(ctx, "insert user err: %v", err)
		return errs.DbError.SetErr(err)
	}

	return nil
}

func (dal *userDal) Update(ctx context.Context, userPO *po.User) errs.Error {
	err := dal.GetGormDB().WithContext(ctx).
		Omit("user_id").
		Where("user_id", userPO.UserID).
		Updates(userPO).Error
	if err != nil {
		logs.CtxErrorf(ctx, "update user err: %v", err)
		return errs.DbError.SetErr(err)
	}

	return nil
}

func (dal *userDal) SelectByID(ctx context.Context, userID string) (*po.User, errs.Error) {
	result := po.NewUser()
	if err := dal.GetGormDB().WithContext(ctx).
		Where("user_id", userID).
		Take(result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logs.CtxErrorf(ctx, "select user by id err: %v", err)
		return nil, errs.DbError.SetErr(err)
	}

	return result, nil
}

func (dal *userDal) SelectByAccount(ctx context.Context, accountID string) (*po.User, errs.Error) {
	result := po.NewUser()
	if err := dal.GetGormDB().WithContext(ctx).
		Where("account_id", accountID).
		Take(result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logs.CtxErrorf(ctx, "select user by account err: %v", err)
		return nil, errs.DbError.SetErr(err)
	}

	return result, nil
}
