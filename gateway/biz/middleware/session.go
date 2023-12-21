package middleware

import (
	"context"
	"net/http"
	"shield/gateway/biz/repos"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/gorilla/securecookie"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"
)

const sessionAccountInfo = "account_info"

func SessionMiddleware() app.HandlerFunc {
	secret, bizErr := repos.GetRandomSecret(
		context.Background(),
		"kaidog_shield_gateway_session_secret",
		string(securecookie.GenerateRandomKey(1024)),
	)
	if bizErr != nil {
		panic(bizErr)
	}

	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte(secret))
	if err != nil {
		hlog.Error("init redis store fail: %v", err)
		panic(err)
	}

	_ = redis.SetKeyPrefix(store, "shield_gateway_")

	store.Options(
		sessions.Options{
			Path:     "/",
			Domain:   "",
			MaxAge:   86400,
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		})

	return sessions.New(sessionAccountInfo, store)
}
