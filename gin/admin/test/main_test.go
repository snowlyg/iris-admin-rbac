package test

import (
	_ "embed"
	"os"
	"testing"

	"github.com/snowlyg/helper/tests"
	rbac "github.com/snowlyg/iris-admin-rbac/gin"
	web_tests "github.com/snowlyg/iris-admin/server/web/tests"
	"github.com/snowlyg/iris-admin/server/web/web_gin"
)

//go:embed mysqlPwd.txt
var mysqlPwd string

//go:embed redisPwd.txt
var redisPwd string

var TestServer *web_gin.WebServer
var TestClient *tests.Client

func TestMain(m *testing.M) {
	var uuid string
	uuid, TestServer = web_tests.BeforeTestMainGin(mysqlPwd, redisPwd, 4, rbac.PartyFunc, rbac.SeedFunc)
	code := m.Run()
	web_tests.AfterTestMain(uuid, rbac.LogoutUrl, TestClient)

	os.Exit(code)
}
