package tests

import (
	"net/http"
	"testing"

	"github.com/snowlyg/helper/str"
	"github.com/snowlyg/httptest"
	rbac "github.com/snowlyg/iris-admin-rbac/iris"
	"github.com/snowlyg/iris-admin/server/web"
)

func TestLoginError(t *testing.T) {
	TestClient := httptest.Instance(t, TestServer.GetEngine(), str.Join("http://", web.CONFIG.System.Addr))
	httptest.Instance(t, TestServer.GetEngine(), str.Join("http://", web.CONFIG.System.Addr))
	TestClient.Login(rbac.LoginUrl, "", httptest.NewResponses(http.StatusBadRequest, "用户名或密码错误", nil), httptest.NewWithJsonParamFunc(map[string]interface{}{"username": "admin1", "password": "123456"}))

}
