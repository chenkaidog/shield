package kitex

import (
	"context"
	"shield/common/logs"
	"time"

	"github.com/cloudwego/kitex/pkg/consts"
	"github.com/cloudwego/kitex/pkg/endpoint"
)

func ServerLogMW(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request, response interface{}) error {
		startTime := time.Now()

		methodName, ok := ctx.Value(consts.CtxKeyMethod).(string)
		if !ok {
			methodName = ""
		}

		var reqBody, respBody interface{} = request, response
		if reqArg, ok := request.(interface{ GetFirstArgument() interface{} }); ok {
			reqBody = reqArg.GetFirstArgument()
		}
		if respArg, ok := response.(interface{ GetResult() interface{} }); ok {
			respBody = respArg.GetResult()
		}

		logs.CtxInfo(ctx, "[%s] request body: %v", methodName, reqBody)
		if err := next(ctx, request, response); err != nil {
			return err
		}
		logs.CtxInfo(ctx, "[%s] resp body: %v, cost: %dms", methodName, respBody, time.Since(startTime)/time.Millisecond)

		return nil
	}
}
