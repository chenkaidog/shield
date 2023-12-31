package kitex

import (
	"context"
	"shield/common/logs"
	"shield/common/utils/sensitive"
	"time"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
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

		methodName := rpcinfo.GetRPCInfo(ctx).Invocation().MethodName()

		if reqArg, ok := request.(interface{ GetFirstArgument() interface{} }); ok {
			reqBody := reqArg.GetFirstArgument()
			logs.CtxInfof(ctx, "[%s] request body: %v", methodName, sensitiveMarshal.SafeMarshal(reqBody))
		}

		err := next(ctx, request, response)
		if err != nil {
			logs.CtxErrorf(ctx, "[%s] request fails: %s", methodName, err.Error())
			return err
		}
		if respArg, ok := response.(interface{ GetResult() interface{} }); ok {
			respBody := respArg.GetResult()
			logs.CtxInfof(ctx, "[%s] resp body: %v, cost: %dms",
				methodName, sensitiveMarshal.SafeMarshal(respBody), time.Since(startTime)/time.Millisecond)
		}

		return nil
	}
}
