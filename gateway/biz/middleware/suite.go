package middleware

import (
	"shield/common/middleware/hertz"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/logger/accesslog"
)

func Suite() []app.HandlerFunc {
	InitLogger()

	var funcList []app.HandlerFunc
	return append(
		funcList,
		accesslog.New(),
		hertz.ServerTraceMW(),
		SessionMiddleware(),
		CsrfMiddleware(),
		RecoveryMiddleware(),
	)
}
