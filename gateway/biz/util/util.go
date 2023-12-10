package util

import "github.com/cloudwego/hertz/pkg/app"

func GetIp(c *app.RequestContext) string {
	return "0.0.0.0"
}

func GetDevice(c *app.RequestContext) string {
	return "unknown"
}