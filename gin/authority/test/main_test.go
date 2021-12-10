package test

import (
	_ "embed"
	"os"
	"testing"

	"github.com/snowlyg/helper/tests"
	rbac "github.com/snowlyg/iris-admin-rbac/gin"
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
	uuid, TestServer = web_gin.BeforeTestMain(mysqlPwd, redisPwd, 3, rbac.PartyFunc, rbac.SeedFunc)
	code := m.Run()
	web_gin.AfterTestMain(uuid, TestClient)

	os.Exit(code)
}
