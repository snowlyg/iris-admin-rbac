package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/helper/str"
	"github.com/snowlyg/httptest"
	rbac "github.com/snowlyg/iris-admin-rbac/gin"
	"github.com/snowlyg/iris-admin/server/web"
	"github.com/snowlyg/iris-admin/server/web/web_gin/response"
)

var (
	url = "/api/v1/public"
)

func TestPublicCaptcha(t *testing.T) {

	TestClient := httptest.Instance(t, TestServer.GetEngine(), str.Join("http://", web.CONFIG.System.Addr))
	TestClient.Login(rbac.LoginUrl, "", httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, rbac.LoginResponse))
	if TestClient == nil {
		return
	}
	pageKeys := httptest.Responses{
		{Key: "captchaId", Value: "", Type: "notempty"},
		{Key: "enable", Value: true},
		{Key: "picPath", Value: "", Type: "notempty"},
	}
	requestParams := map[string]interface{}{"page": 1, "pageSize": 10, "orderBy": "id"}
	TestClient.GET(fmt.Sprintf("%s/captcha", url), httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, pageKeys), httptest.NewWithQueryObjectParamFunc(requestParams))
}

func TestLoginError(t *testing.T) {
	TestClient := httptest.Instance(t, TestServer.GetEngine(), str.Join("http://", web.CONFIG.System.Addr))
	requestParams := map[string]interface{}{"username": "admin1", "password": "123456"}
	TestClient.POST(fmt.Sprintf("%s/admin/login", url), httptest.NewResponses(http.StatusBadRequest, "用户名或密码错误", nil), httptest.NewWithJsonParamFunc(requestParams))
}
