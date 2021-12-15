package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/iris-admin-rbac/gin/admin"
	"github.com/snowlyg/iris-admin-rbac/gin/api"
	"github.com/snowlyg/iris-admin-rbac/gin/authority"
	"github.com/snowlyg/iris-admin-rbac/gin/public"
	"github.com/snowlyg/iris-admin/migration"
	"github.com/snowlyg/iris-admin/server/operation"
	"github.com/snowlyg/iris-admin/server/web/web_gin"
	"github.com/snowlyg/multi"
)

// Party v1 模块
func Party(group *gin.RouterGroup) {
	api.Group(group)
	admin.Group(group)
	authority.Group(group)
	public.Group(group)
}

var LoginUrl = "/api/v1/public/admin/login"
var LogoutUrl = "/api/v1/public/logout"
var LoginResponse = httptest.Responses{
	{Key: "status", Value: http.StatusOK},
	{Key: "message", Value: "操作成功"},
	{Key: "data",
		Value: httptest.Responses{
			{Key: "accessToken", Value: "", Type: "notempty"},
		},
	},
}
var LogoutResponse = httptest.Responses{
	{Key: "status", Value: http.StatusOK},
	{Key: "message", Value: "操作成功"},
}

// 加载模块
var PartyFunc = func(wi *web_gin.WebServer) {
	Party(wi.GetRouterGroup("/api/v1"))
}

//  填充数据
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
