// Code generated by hertz generator.

package gateway

import (
	"context"
	"net/http"
	"shield/gateway/biz/model/consts"
	"shield/gateway/biz/repos"
	"shield/gateway/biz/util"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

func rootMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		func(ctx context.Context, c *app.RequestContext) {
			sess := sessions.Default(c)
			accountId, ok := sess.Get(consts.SessionAccountId).(string)
			if ok {
				c.Set(consts.ContextAccountId, accountId)
			}

			c.Next(ctx)
		},
	}
}

func _createaccountMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _createuserMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _loginMw() []app.HandlerFunc {
	return nil
}

func _queryloginrecordMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _logoutMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		func(ctx context.Context, c *app.RequestContext) {
			_, ok := util.GetAccountId(c)
			if !ok {
				c.AbortWithMsg("user not login", http.StatusUnauthorized)
				return
			}

			c.Next(ctx)
		},
	}
}

func _resetpasswordMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _switchaccountstatusMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updatepasswordMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _updateuserinfoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _queryuserinfoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _adminMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _userMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _queryaccountMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _operatorMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		func(ctx context.Context, c *app.RequestContext) {
			accountId, ok := util.GetAccountId(c)
			if !ok {
				c.AbortWithMsg("user not login", http.StatusUnauthorized)
				// todo: redirect
				return
			}
			sessID, bizErr := repos.GetAccountSessionID(ctx, accountId)
			if bizErr != nil {
				c.AbortWithMsg("server error", http.StatusInternalServerError)
				return
			}

			if sessID != sessions.Default(c).ID() {
				c.AbortWithMsg("login timeout", http.StatusUnauthorized)
				// todo: redirect
				return
			}

			c.Next(ctx)
		},
	}
}

func _queryloginrecord0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _queryuserinfo0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _queryselfloginrecordMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _queryselfuserinfoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _accountMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _account_statusMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _login_recordMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _passwordMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _user_infoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _login_record0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _password0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _user_info0Mw() []app.HandlerFunc {
	// your code...
	return nil
}
