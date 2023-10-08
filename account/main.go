package main

import (
	"log"
	account "shield/account/kitex_gen/kaidog/shield/account/accountservice"
)

func main() {
	svr := account.NewServer(new(AccountServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
