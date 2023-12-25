package repos

import (
	"context"
	"shield/common/errs"
	"shield/gateway/biz/repos/redis"
)

func SetAccountCsrfToken(ctx context.Context) {

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
