package tests

import (
	"os"
	"testing"

	"github.com/snowlyg/httptest"
	rbac "github.com/snowlyg/iris-admin-rbac/iris"
	"github.com/snowlyg/iris-admin/server/web/common"
	"github.com/snowlyg/iris-admin/server/web/web_iris"
)

var TestServer *web_iris.WebServer
var TestClient *httptest.Client

func TestMain(m *testing.M) {

	var uuid string
	uuid, TestServer = common.BeforeTestMainIris(rbac.PartyFunc, rbac.SeedFunc)
	code := m.Run()
	common.AfterTestMain(uuid, true)

	os.Exit(code)
}
