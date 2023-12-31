package repos

import (
	"context"
	"fmt"
	"shield/common/errs"
	"shield/gateway/biz/repos/redis"
)

func genAccountSessionIDKey(accountId string) string {
	return fmt.Sprintf("shield_session_id_%s", accountId)
}

func SetAccountSessionID(ctx context.Context, accountId, token string) errs.Error {
	return redis.Set(ctx, genAccountSessionIDKey(accountId), token, 0)
}

func GetAccountSessionID(ctx context.Context, accountId string) (string, errs.Error) {
	return redis.Get(ctx, genAccountSessionIDKey(accountId))
}

func RemoveAccountSessionID(ctx context.Context, accountId string) errs.Error {
	return redis.Del(ctx, genAccountSessionIDKey(accountId))
}

func GetRandomSecret(ctx context.Context, key, secret string) (string, errs.Error) {
	ok, err := redis.SetNX(ctx, key, secret, 0)
	if err != nil {
		return "", err
	}
	if ok {
		return secret, nil
	}

	existedSecret, err := redis.Get(ctx, key)
	if err != nil {
		return "", err
	}

	return existedSecret, nil
}
