package tests

import (
	"os"
	"testing"

	rbac "github.com/snowlyg/iris-admin-rbac/gin"
	"github.com/snowlyg/iris-admin/server/web/common"
	"github.com/snowlyg/iris-admin/server/web/web_gin"
)

var TestServer *web_gin.WebServer

func TestMain(m *testing.M) {

	var uuid string
	uuid, TestServer = common.BeforeTestMainGin(rbac.PartyFunc, rbac.SeedFunc)
	code := m.Run()
	common.AfterTestMain(uuid, true)

	os.Exit(code)
}
