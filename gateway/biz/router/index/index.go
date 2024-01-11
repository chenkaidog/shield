package index

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

func Login(ctx context.Context, c *app.RequestContext) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func Home(ctx context.Context, c *app.RequestContext) {
	c.HTML(http.StatusOK, "home.html", nil)
}

func Error(ctx context.Context, c *app.RequestContext) {
	errCode := c.Param("error_code")

	c.HTML(http.StatusOK, "error.html", map[string]interface{}{
		"error_code": errCode,
		"msg":        "something wrong",
	})
}
