package kitex

import (
	"context"
	"shield/common/errs"
	"shield/common/logs"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

func PanicRecoverMW(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, args, result interface{}) (err error) {
		if err := next(ctx, args, result); err != nil {
			ri := rpcinfo.GetRPCInfo(ctx)
			if stats := ri.Stats(); stats != nil {
				if panicked, err := stats.Panicked(); panicked {
					// `err` 就是框架调用 recover() 收到的对象
					if detailErr, ok := err.(*kerrors.DetailedError); ok {
						logs.CtxError(ctx, "panic occur: %s\n%s", detailErr.Error(), detailErr.Stack())
					} else {
						logs.CtxError(ctx, "panic occur: %v", err)
					}

					return errs.ServerError.SetMsg("internal server error")
				}
			}

			return err
		}

		return nil
	}
}
