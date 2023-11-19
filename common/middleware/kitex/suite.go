package kitex

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/server"
)

type ServerSuite struct{}

func (*ServerSuite) Options() []server.Option {
	return []server.Option{
		server.WithMiddleware(ServerTraceMW),
		server.WithMiddleware(KitexLogMW),
		server.WithMiddleware(ServerErrorHandlerMW),
		server.WithMiddleware(RequestValidatorMW),
		server.WithMiddleware(PanicRecoverMW),
	}
}

func NewServerSuite() *ServerSuite {
	return new(ServerSuite)
}

type ClientSuite struct{}

func (*ClientSuite) Options() []client.Option {
	return []client.Option{
		client.WithMiddleware(ClientTraceMW),
		client.WithMiddleware(KitexLogMW),
		client.WithMiddleware(ClientErrorHandlerMW),
		client.WithMiddleware(RequestValidatorMW),
		client.WithMiddleware(PanicRecoverMW),
	}
}

func NewClientSuite() *ClientSuite {
	return new(ClientSuite)
}
