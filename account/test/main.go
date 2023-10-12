package main

import (
	"context"
	"fmt"
	"shield/account/kitex_gen/base"
	"shield/account/kitex_gen/kaidog/shield/account"
	"shield/account/kitex_gen/kaidog/shield/account/accountservice"

	"github.com/cloudwego/kitex/client"
)

func main() {
	ctx := context.Background()
	client := accountservice.MustNewClient("AccountService", client.WithHostPorts("0.0.0.0:8888"))
	resp, err := client.Login(ctx, &account.AccountLoginReq{
		Username: "bobbobobo",
		Password: "12345678",
		Ipv4:     "127.0.0.1",
		Device:   "PC",
		Base: &base.BaseReq{
			TraceID: "00000000000000",
			LogID:   "",
			SpanID:  "",
		},
	})
	fmt.Println(resp, err)
}
