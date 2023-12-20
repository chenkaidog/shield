package middleware

import (
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/gorilla/securecookie"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"
)

const sessionAccountInfo = "account_info"

func SessionMiddleware() app.HandlerFunc {
	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", securecookie.GenerateRandomKey(1024))
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
