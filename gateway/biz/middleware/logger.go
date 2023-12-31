package middleware

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"shield/common/middleware/hertz"
)

func InitLogger() {
	hlog.SetLogger(hertz.NewHertzLogger())
}
