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

func UpdateAccount(ctx context.Context,account *po.Account ) errs.Error {
	return mysql.NewAccountDal().Update(ctx, account)
}

func SelectAccountByID(ctx context.Context, accountID string)(*po.Account, errs.Error)  {
	return mysql.NewAccountDal().SelectByID(ctx, accountID)
}

func SelectAccountByUsername(ctx context.Context, username string) (*po.Account, errs.Error) {
	return mysql.NewAccountDal().SelectByUsername(ctx, username)
}