package iris

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/iris-admin-rbac/iris/auth"
	"github.com/snowlyg/iris-admin-rbac/iris/file"
	"github.com/snowlyg/iris-admin-rbac/iris/oplog"
	"github.com/snowlyg/iris-admin-rbac/iris/perm"
	"github.com/snowlyg/iris-admin-rbac/iris/role"
	"github.com/snowlyg/iris-admin-rbac/iris/user"
)

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
