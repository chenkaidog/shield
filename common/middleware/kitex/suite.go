package kitex

import "github.com/cloudwego/kitex/server"

type ServerSuite struct{}

func (*ServerSuite) Options() []server.Option {
	return []server.Option{
		server.WithMiddleware(ServerTraceMW),
		server.WithMiddleware(ServerLogMW),
		server.WithMiddleware(ErrorHandlerMW),
		server.WithMiddleware(RequestValidatorMW),
		server.WithMiddleware(PanicRecoverMW),
	}
}

func NewServerSuite() *ServerSuite {
	return new(ServerSuite)
}
