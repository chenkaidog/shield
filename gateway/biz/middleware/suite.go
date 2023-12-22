package middleware

import (
	"shield/common/middleware/hertz"

	"github.com/cloudwego/hertz/pkg/app"
)

func Suite() []app.HandlerFunc {
	InitLogger()

	var funcList []app.HandlerFunc
	return append(
		funcList,
		hertz.ServerTraceMW(),
		SessionMiddleware(),
		CsrfMiddleware(),
		RecoveryMiddleware(),
	)
}
