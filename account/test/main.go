package main

import (
	"context"
	"fmt"
	"shield/account/kitex_gen/kaidog/shield/account"
	"shield/account/kitex_gen/kaidog/shield/account/accountservice"
)

func main() {
	ctx := context.Background()
	client := accountservice.MustNewClient("kaidog.shield.account")
	resp, err := client.Login(ctx, &account.AccountLoginReq{
		Username: "bob",
		Password: "123456",
	})
	fmt.Println(resp, err)
}
