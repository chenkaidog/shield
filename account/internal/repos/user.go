package repos

import (
	"context"
	"shield/account/internal/model/po"
	"shield/account/internal/repos/mysql"
	"shield/common/errs"
)

func CreateUser(ctx context.Context, user *po.User) errs.Error {
	return mysql.NewUserDal().Insert(ctx, user)
}

func UpdateUser(ctx context.Context, user *po.User) errs.Error {
	return mysql.NewUserDal().Update(ctx, user)
}

func SelectUserByID(ctx context.Context, userID string) (*po.User, errs.Error) {
	return mysql.NewUserDal().SelectByID(ctx, userID)
}

func SelectUserByAccountID(ctx context.Context, accountID string) (*po.User, errs.Error) {
	return mysql.NewUserDal().SelectByAccount(ctx, accountID)
}
