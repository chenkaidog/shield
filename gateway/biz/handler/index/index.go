package index

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

func Login(ctx context.Context, c *app.RequestContext) {
	c.HTML(http.StatusOK, "login.html", nil)
}