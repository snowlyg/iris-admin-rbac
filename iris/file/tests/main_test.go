package tests

import (
	_ "embed"
	"os"
	"testing"

	"github.com/snowlyg/helper/tests"
	rbac "github.com/snowlyg/iris-admin-rbac/iris"
	web_tests "github.com/snowlyg/iris-admin/server/web/tests"
	"github.com/snowlyg/iris-admin/server/web/web_iris"
)

//go:embed mysqlPwd.txt
var mysqlPwd string

//go:embed redisPwd.txt
var redisPwd string

var TestServer *web_iris.WebServer
var TestClient *tests.Client

func TestMain(m *testing.M) {
	var uuid string
	uuid, TestServer = web_tests.BeforeTestMainIris(mysqlPwd, redisPwd, 1, rbac.PartyFunc, rbac.SeedFunc)
	code := m.Run()
	web_tests.AfterTestMain(uuid, rbac.LogoutUrl, TestClient)

	os.Exit(code)
}
