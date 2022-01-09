package test

import (
	"os"
	"testing"

	"github.com/snowlyg/httptest"
	rbac "github.com/snowlyg/iris-admin-rbac/gin"
	"github.com/snowlyg/iris-admin/server/web/common"
	"github.com/snowlyg/iris-admin/server/web/web_gin"
)

var TestServer *web_gin.WebServer
var TestClient *httptest.Client

func TestMain(m *testing.M) {

	var uuid string
	uuid, TestServer = common.BeforeTestMainGin(rbac.PartyFunc, rbac.SeedFunc)
	code := m.Run()
	common.AfterTestMain(uuid, true)

	os.Exit(code)
}
