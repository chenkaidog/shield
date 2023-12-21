package middleware

import (
	"context"
	"shield/gateway/biz/repos"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gorilla/securecookie"
	"github.com/hertz-contrib/csrf"
)

func CsrfMiddleware() app.HandlerFunc {
	secret, err := repos.GetRandomSecret(
		context.Background(),
		"kaidog_shield_gateway_csrf_secret",
		string(securecookie.GenerateRandomKey(1024)),
	)
	if err != nil {
		panic(err)
	}
	return csrf.New(
		csrf.WithSecret(secret),
		csrf.WithNext(func(ctx context.Context, c *app.RequestContext) bool {
			switch string(c.Path()) {
			case "/login":
				return true
			default:
				return false
			}
		}),
	)
}
