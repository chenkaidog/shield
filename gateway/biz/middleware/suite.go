package middleware

import (
	"github.com/cloudwego/hertz/pkg/app"
	"shield/common/middleware/hertz"
)

func Suite() []app.HandlerFunc {
	var funcList []app.HandlerFunc
	return append(
		funcList,
		hertz.ServerTraceMW(),
		SessionMiddleware(),
		CsrfMiddleware(),
		RecoveryMiddleware(),
	)
}