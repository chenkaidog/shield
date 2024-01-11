package util

import (
	"net/http"
	"reflect"
	"shield/common/errs"

	"github.com/cloudwego/hertz/pkg/app"
)

const (
	fieldSuccess = "Success"
	fieldCode    = "Code"
	fieldMsg     = "Msg"
)

// BuildBizResp builds resp code to resp.
// resp must be a pointer, or can not reflect to its fields.
func BuildBizResp(c *app.RequestContext, resp interface{}, err errs.Error) {
	if err == nil {
		fillRespCode(resp, errs.Success)
	} else {
		fillRespCode(resp, err)
	}

	c.JSON(http.StatusOK, resp)
}

func fillRespCode(resp interface{}, err errs.Error) {
	respVal := reflect.ValueOf(resp)

	if respVal.Kind() == reflect.Ptr && !respVal.IsNil() {
		respVal = respVal.Elem()
		if respVal.Kind() == reflect.Struct {
			if _, ok := respVal.Type().FieldByName(fieldSuccess); ok {
				if success := respVal.FieldByName(fieldSuccess); success.CanSet() {
					success.SetBool(errs.ErrorEqual(err, errs.Success))
				}
			}
			if _, ok := respVal.Type().FieldByName(fieldCode); ok {
				if code := respVal.FieldByName(fieldCode); code.CanSet() {
					code.SetInt(int64(err.Code()))
				}
			}
			if _, ok := respVal.Type().FieldByName(fieldMsg); ok {
				if msg := respVal.FieldByName(fieldMsg); msg.CanSet() {
					msg.SetString(err.Msg())
				}
			}
		}
	}
}
