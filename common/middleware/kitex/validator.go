package kitex

import (
	"context"
	"shield/common/errs"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

func RequestValidatorMW(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, args, result interface{}) error {
		if gfa, ok := args.(interface{ GetFirstArgument() interface{} }); ok {
			req := gfa.GetFirstArgument()
			if rv, ok := req.(interface{ IsValid() error }); ok {
				if err := rv.IsValid(); err != nil {
					return errs.ParamError.SetErr(err)
				}
			}
		}

		if err := next(ctx, args, result); err != nil {
			return err
		}

		if gr, ok := result.(interface{ GetResult() interface{} }); ok {
			resp := gr.GetResult()
			if rv, ok := resp.(interface{ IsValid() error }); ok {
				if err := rv.IsValid(); err != nil {
					return errs.ParamError.SetErr(err)
				}
			}
		}

		return nil
	}
}
