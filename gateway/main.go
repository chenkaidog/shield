// Code generated by hertz generator.

package main

import (
	"shield/gateway/biz/rpc"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	rpc.Init()

	h := server.Default()

	register(h)
	h.Spin()
}
