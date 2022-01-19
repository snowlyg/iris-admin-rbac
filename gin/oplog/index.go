package oplog

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/iris-admin-rbac/gin/middleware"
)

// Party 操作日志
func Group(group *gin.RouterGroup) {
	router := group.Group("/oplog", middleware.Auth(), middleware.CasbinHandler())
	{
		router.GET("/getOplogList", GetOplogList) // 获取操作日志列表
	}
}
