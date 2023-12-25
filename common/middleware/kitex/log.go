package kitex

import (
	"context"
	"shield/common/logs"
	"shield/common/utils/sensitive"
	"time"

	"github.com/cloudwego/kitex/pkg/consts"
	"github.com/cloudwego/kitex/pkg/endpoint"
)

func init() {
	sensitiveMarshal = sensitive.NewSensitiveMarshal("password")
}

var sensitiveMarshal *sensitive.SensitiveMarshal

func SetSensitiveWord(words ...string) {
	sensitiveMarshal.AddSensitiveWord(words...)
}

func KitexLogMW(next endpoint.Endpoint) endpoint.Endpoint {
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

		logs.CtxInfof(ctx, "[%s] request body: %v", methodName, sensitiveMarshal.SafeMarshal(reqBody))
		defer func() {
			logs.CtxInfof(ctx, "[%s] resp body: %v, cost: %dms",
				methodName, sensitiveMarshal.SafeMarshal(respBody), time.Since(startTime)/time.Millisecond)
		}()
		if err := next(ctx, request, response); err != nil {
			return err
		}

		return nil
	}
}
