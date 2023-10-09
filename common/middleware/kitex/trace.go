package kitex

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
)

func TraceMW(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, args, result interface{}) error {

		if err := next(ctx, args, result); err != nil {
			return err
		}

		return nil
	}
}
