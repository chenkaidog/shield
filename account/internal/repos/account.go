package repos

import (
	"context"
	"shield/account/internal/repos/mysql"
	"shield/account/model/po"
	"shield/common/errs"
)

func CreateAccount(ctx context.Context, account *po.Account) errs.Error {
	return mysql.NewAccountDal().Insert(ctx, account)
}

func UpdateAccount(ctx context.Context, account *po.Account) errs.Error {
	return mysql.NewAccountDal().Update(ctx, account)
}

func SelectAccountByID(ctx context.Context, accountId string) (*po.Account, errs.Error) {
	return mysql.NewAccountDal().SelectByID(ctx, accountId)
}

func SelectAccount(ctx context.Context, limit, offset int) ([]*po.Account, int64, errs.Error) {
	return mysql.NewAccountDal().Select(ctx, limit, offset)
}

func SelectAccountByUsername(ctx context.Context, username string) (*po.Account, errs.Error) {
	return mysql.NewAccountDal().SelectByUsername(ctx, username)
}
