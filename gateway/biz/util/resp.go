package util

import (
	"net/http"
	"shield/common/errs"
	"shield/gateway/biz/model/resp"

	"github.com/cloudwego/hertz/pkg/app"
)

func BuildRespParamErr(c *app.RequestContext, err error) {
	c.JSON(
		http.StatusOK,
		resp.NewFailResp(
			errs.ParamError.Code(),
			err.Error(),
		),
	)
}

func BuildRespBizErr(c *app.RequestContext, err errs.Error) {
	c.JSON(
		http.StatusOK,
		resp.NewFailResp(
			err.Code(),
			err.Msg(),
		),
	)
}

func BuildRespSuccess(c *app.RequestContext, body interface{}) {
	c.JSON(http.StatusOK, resp.NewSuccessResp(body))
}