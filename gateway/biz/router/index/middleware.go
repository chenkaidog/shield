package index

import (
	"bytes"
	"context"
	"net/http"
	"net/url"
	"shield/gateway/biz/model/consts"
	"shield/gateway/biz/repos"
	"strconv"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

func IndexMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		ErrorHandleMW(),
		AuthMiddleware(),
	}
}

func ErrorHandleMW() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		c.Next(ctx)

		switch statusCode := c.Response.StatusCode(); statusCode {
		case http.StatusOK:
			return
		case http.StatusUnauthorized:
			if bytes.Equal(c.Path(), []byte("/login")) {
				return
			}

			sb := strings.Builder{}
			sb.WriteString("/login?redirect=")
			sb.WriteString(url.QueryEscape(c.URI().String()))
			c.Redirect(http.StatusOK, []byte(sb.String()))
		default:
			sb := strings.Builder{}
			sb.WriteString("/error/")
			sb.WriteString(strconv.Itoa(statusCode))
			c.Redirect(statusCode, []byte(sb.String()))
		}
	}
}

func AuthMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		sess := sessions.Default(c)
		accountId, ok := sess.Get(consts.SessionAccountId).(string)
		if !ok {
			c.AbortWithMsg("user not login", http.StatusUnauthorized)
			return
		}
		sessID, bizErr := repos.GetAccountSessionID(ctx, accountId)
		if bizErr != nil {
			c.AbortWithMsg("server error", http.StatusInternalServerError)
			return
		}

		if sessID != sessions.Default(c).ID() {
			c.AbortWithMsg("user not login", http.StatusUnauthorized)
			return
		}

		c.Next(ctx)
	}
}
