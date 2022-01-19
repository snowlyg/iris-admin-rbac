package oplog

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/iris-admin-rbac/iris/middleware"
)

// Party 操作日志
func Party() func(index iris.Party) {
	return func(index iris.Party) {
		index.Use(middleware.MultiHandler(), middleware.Casbin())
		index.Get("/", GetAll).Name = "操作日志列表"
	}
}
