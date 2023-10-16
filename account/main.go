package main

import (
	"log"
	"shield/account/internal/config"
	"shield/account/internal/repos"
	account "shield/account/kitex_gen/kaidog/shield/account/accountservice"
	"shield/common/middleware/kitex"

	"github.com/cloudwego/kitex/server"
)

func main() {
	config.Init()
	repos.Init()

	svr := account.NewServer(
		new(AccountServiceImpl),
		server.WithSuite(kitex.NewServerSuite()),
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
