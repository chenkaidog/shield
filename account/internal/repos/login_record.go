package repos

import (
	"context"
	"shield/account/internal/model/po"
	"shield/account/internal/repos/mysql"
	"shield/common/errs"
)

func CreateLoginRecord(ctx context.Context, loginRecordPO *po.LoginRecord) errs.Error {
	return mysql.NewLoginRecordDal().Insert(ctx, loginRecordPO)
}

func SelectLoginRecord(ctx context.Context, accountID string, limit, offset int) ([]*po.LoginRecord, int64, errs.Error) {
	return mysql.NewLoginRecordDal().Select(ctx, accountID, limit, offset)
}
