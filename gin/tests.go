package v1

import (
	"os"

	"github.com/snowlyg/helper/str"
	"github.com/snowlyg/httptest"
	"github.com/snowlyg/iris-admin-rbac/gin/admin"
	"github.com/snowlyg/iris-admin-rbac/gin/api"
	"github.com/snowlyg/iris-admin-rbac/gin/authority"
	"github.com/snowlyg/iris-admin/migration"
	"github.com/snowlyg/iris-admin/server/cache"
	"github.com/snowlyg/iris-admin/server/operation"
	"github.com/snowlyg/iris-admin/server/web/web_gin"
	"github.com/snowlyg/iris-admin/server/zap_server"
	"github.com/snowlyg/multi"
)

var prefix = "/api/v1"
var LoginUrl = str.Join(prefix, "/public/admin/login")
var LogoutUrl = str.Join(prefix, "/public/logout")
var LoginResponse = httptest.Responses{
	{Key: "accessToken", Value: "", Type: "notempty"},
	{Key: "user", Value: httptest.Responses{
		{Key: "id", Value: 0, Type: "ge"}}},
}

// 加载模块
var PartyFunc = func(wi *web_gin.WebServer) {
	driverType := os.Getenv("driverType")
	if driverType == "" {
		driverType = "jwt"
	}
	config := &multi.Config{DriverType: driverType, HmacSecret: nil}
	if driverType == "redis" {
		config.UniversalClient = cache.Instance()
	}
	// 初始化驱动
	err := multi.InitDriver(config)
	if err != nil {
		zap_server.ZAPLOG.Panic("err")
	}
	Party(wi.GetRouterGroup(prefix))
}

// 填充数据
var SeedFunc = func(wi *web_gin.WebServer, mc *migration.MigrationCmd) {
	mc.AddMigration(api.GetMigration(), authority.GetMigration(), admin.GetMigration(), operation.GetMigration())
	routes, _ := wi.GetSources()
	// 权鉴模块全部为管理员权限
	authorityTypes := map[string]int{}
	for _, route := range routes {
		authorityTypes[route["path"]] = multi.AdminAuthority
	}
	// notice : 注意模块顺序
	mc.AddSeed(api.New(routes, authorityTypes), authority.Source, admin.Source)
}
