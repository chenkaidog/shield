package main

import (
	"log"
	account "shield/account/kitex_gen/kaidog/shield/account/accountservice"
	"shield/common/middleware/kitex"

	"github.com/cloudwego/kitex/server"
)

func main() {
	svr := account.NewServer(
		new(AccountServiceImpl),
		server.WithSuite(kitex.NewServerSuite()),
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
