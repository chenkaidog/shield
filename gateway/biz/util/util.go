package util

import (
	"fmt"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/mssola/user_agent"
)

func GetIp(c *app.RequestContext) string {
	host := string(c.Host())
	idx := strings.Index(host, ":")
	if idx > 0 {
		return host[:idx]
	}

	return host
}

func GetDevice(c *app.RequestContext) string {
	userAgent := user_agent.New(string(c.UserAgent()))
	name, version :=  userAgent.Browser()
	return fmt.Sprintf("%s/%s", name, version)
}