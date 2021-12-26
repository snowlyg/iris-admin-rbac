package iris

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/iris-admin-rbac/iris/auth"
	"github.com/snowlyg/iris-admin-rbac/iris/file"
	"github.com/snowlyg/iris-admin-rbac/iris/oplog"
	"github.com/snowlyg/iris-admin-rbac/iris/perm"
	"github.com/snowlyg/iris-admin-rbac/iris/role"
	"github.com/snowlyg/iris-admin-rbac/iris/user"
	"github.com/snowlyg/iris-admin/migration"
	"github.com/snowlyg/iris-admin/server/operation"
	"github.com/snowlyg/iris-admin/server/web/web_iris"
	"github.com/snowlyg/iris-admin/server/zap_server"
	"github.com/snowlyg/multi"
)

var LoginUrl = "/api/v1/auth/login"
var LogoutUrl = "/api/v1/users/logout"

// Party v1 模块
func Party() func(rbac iris.Party) {
	return func(rbac iris.Party) {
		rbac.PartyFunc("/users", user.Party())
		rbac.PartyFunc("/roles", role.Party())
		rbac.PartyFunc("/perms", perm.Party())
		rbac.PartyFunc("/file", file.Party())
		rbac.PartyFunc("/auth", auth.Party())
		rbac.PartyFunc("/oplog", oplog.Party())
	}
}

// 加载模块
var PartyFunc = func(wi *web_iris.WebServer) {
	// 初始化驱动
	err := multi.InitDriver(&multi.Config{DriverType: "jwt", HmacSecret: nil})
	if err != nil {
		zap_server.ZAPLOG.Panic("err")
	}
	wi.AddModule(web_iris.Party{
		Perfix:    "/api/v1",
		PartyFunc: Party(),
	})
}

//  填充数据
var SeedFunc = func(wi *web_iris.WebServer, mc *migration.MigrationCmd) {
	mc.AddMigration(perm.GetMigration(), role.GetMigration(), user.GetMigration(), operation.GetMigration())
	routes, _ := wi.GetSources()
	// notice : 注意模块顺序
	mc.AddSeed(perm.New(routes), role.Source, user.Source)
}
