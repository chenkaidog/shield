package kitex

import "github.com/cloudwego/kitex/server"

type Suite struct{}

func (*Suite) Options() []server.Option {
	return []server.Option{
		server.WithMiddleware(TraceMW),
		server.WithMiddleware(ServerLogMW),
		server.WithMiddleware(ErrorHandlerMW),
		server.WithMiddleware(RequestValidatorMW),
		server.WithMiddleware(PanicRecoverMW),
	}
}

func NewSuite() *Suite {
	return new(Suite)
}
