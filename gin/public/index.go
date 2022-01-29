package public

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/iris-admin-rbac/gin/middleware"
)

// Group 认证模块
func Group(group *gin.RouterGroup) {
	group.GET("/public/captcha", Captcha)
	group.POST("/public/admin/login", AdminLogin)
	publicRouter := group.Group("/public")
	{
		publicRouter.Use(middleware.Auth(), middleware.CasbinHandler())
		publicRouter.GET("/logout", Logout) // 退出
		publicRouter.GET("/clean", Clear)   //清空授权
	}
}
