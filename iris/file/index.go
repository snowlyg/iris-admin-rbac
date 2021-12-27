package file

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/iris-admin-rbac/iris/middleware"
	"github.com/snowlyg/iris-admin/server/web"
)

// Party 上传文件模块
func Party() func(index iris.Party) {
	return func(index iris.Party) {
		index.Use(middleware.MultiHandler(), middleware.OperationRecord(), middleware.Casbin())
		index.Post("/", iris.LimitRequestBodySize(web.CONFIG.MaxSize+1<<20), Upload).Name = "上传文件"
	}
}
