package util

import (
	"encoding/json"
	"shield/gateway/biz/model/kaidog/shield/gateway"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
)

func TestBuildBizResp(t *testing.T) {
	var resp gateway.LoginResp

	BuildBizResp(app.NewContext(1234), &resp, nil)

	data, _ := json.Marshal(resp)
	t.Logf("%s", data)
}
