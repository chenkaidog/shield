package util

import (
	"shield/gateway/biz/model/consts"
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
	if name, _ := userAgent.Browser(); name != "" {
		return name
	}

	return "UNKNOWN"
}

func GetAccountId(c *app.RequestContext) (string, bool) {
	val, ok := c.Get(consts.ContextAccountId)
	if !ok {
		return "", false
	}

	accId, ok := val.(string)
	if !ok {
		return "", false
	}

	return accId, true
}
