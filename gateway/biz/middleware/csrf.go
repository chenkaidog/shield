package middleware

import (
	"context"
	"net/http"
	"shield/common/logs"
	"shield/common/middleware/hertz/csrf"
	"shield/common/utils/random"
	"shield/gateway/biz/repos"

	"github.com/cloudwego/hertz/pkg/app"
)

func CsrfMiddleware() app.HandlerFunc {
	secret, err := repos.GetRandomSecret(
		context.Background(),
		"kaidog_shield_gateway_csrf_secret",
		random.RandStr(1024),
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
		csrf.WithErrorFunc(func(ctx context.Context, c *app.RequestContext) {
			csrfErr := c.Errors.Last()
			logs.CtxErrorf(ctx, "scrf recover err: %s", csrfErr.Error())
			c.AbortWithMsg(csrfErr.Error(), http.StatusForbidden)
		}),
	)
}
