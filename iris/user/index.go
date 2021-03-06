package user

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/iris-admin-rbac/iris/middleware"
)

// Party 用户
func Party() func(index iris.Party) {
	return func(index iris.Party) {
		index.Use(middleware.MultiHandler(), middleware.Casbin())
		index.Get("/", GetAll).Name = "用户列表"
		index.Get("/{id:uint}", GetUser).Name = "用户详情"
		index.Post("/", CreateUser).Name = "创建用户"
		index.Post("/{id:uint}", UpdateUser).Name = "编辑用户"
		index.Delete("/{id:uint}", DeleteUser).Name = "删除用户"
		index.Get("/logout", Logout).Name = "退出"
		index.Get("/clear", Clear).Name = "清空 token"
		index.Get("/profile", Profile).Name = "个人信息"
		index.Post("/changeAvatar", ChangeAvatar).Name = "修改头像"
	}
}
