// Code generated by Kitex v0.7.2. DO NOT EDIT.
package accountservice

import (
	server "github.com/cloudwego/kitex/server"
	account "shield/gateway/kitex_gen/kaidog/shield/account"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler account.AccountService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
