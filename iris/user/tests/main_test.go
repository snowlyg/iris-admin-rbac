package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/snowlyg/helper/tests"
	rbac "github.com/snowlyg/iris-admin-rbac/iris"
	web_tests "github.com/snowlyg/iris-admin/server/web/tests"
	"github.com/snowlyg/iris-admin/server/web/web_iris"
)

var TestServer *web_iris.WebServer
var TestClient *tests.Client

func TestMain(m *testing.M) {
	mysqlPwd := os.Getenv("mysqlPwd")
	redisPwd := os.Getenv("redisPwd")
	fmt.Println(mysqlPwd, redisPwd)
	var uuid string
	uuid, TestServer = web_tests.BeforeTestMainIris(mysqlPwd, redisPwd, 4, rbac.PartyFunc, rbac.SeedFunc)
	code := m.Run()
	web_tests.AfterTestMain(uuid)

	os.Exit(code)
}
