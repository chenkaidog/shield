package redis

import (
	"context"
	"fmt"
	"shield/common/errs"
	hook "shield/common/middleware/redis"
	"shield/gateway/biz/config"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdbClient *redis.Client

func InitReis() {
	rdbClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.GetRedisConf().IP, config.GetRedisConf().Port),
		Password: config.GetRedisConf().Password,
		DB:       config.GetRedisConf().DB,
	})

	rdbClient.AddHook(hook.NewRedisHook())
}

func SetNX(ctx context.Context, key, value string, timeout time.Duration) (bool, errs.Error) {
	ok, err := rdbClient.SetNX(ctx, key, value, timeout).Result()
	if err != nil {
		return false, errs.RedisError.SetErr(err)
	}

	return ok, nil
}

func Set(ctx context.Context, key, value string, timeout time.Duration) errs.Error {
	if err := rdbClient.Set(ctx, key, value, timeout).Err(); err != nil {
		return errs.RedisError.SetErr(err)
	}

	return nil
}

func Get(ctx context.Context, key string) (string, errs.Error) {
	result, err := rdbClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", nil
		}

		return "", errs.RedisError.SetErr(err)
	}

	return result, nil
}
