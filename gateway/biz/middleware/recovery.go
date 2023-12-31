package middleware

import (
	"context"
	"net/http"
	"shield/common/logs"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
)

func MyRecoveryHandler(c context.Context, ctx *app.RequestContext, err interface{}, stack []byte) {
	logs.CtxErrorf(c, "[Recovery] err=%v\nstack=%s", err, stack)
	ctx.AbortWithStatus(http.StatusInternalServerError)
}

func RecoveryMiddleware() app.HandlerFunc {
	return recovery.Recovery(recovery.WithRecoveryHandler(MyRecoveryHandler))
}
