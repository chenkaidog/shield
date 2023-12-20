package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gorilla/securecookie"
	"github.com/hertz-contrib/csrf"
)

func CsrfMiddleware() app.HandlerFunc {
	return csrf.New(
		csrf.WithSecret(string(securecookie.GenerateRandomKey(1024))),
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
