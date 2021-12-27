package api

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/iris-admin-rbac/gin/middleware"
)

func Group(group *gin.RouterGroup) {
	apiRouter := group.Group("/api", middleware.Auth(), middleware.CasbinHandler())
	{
		apiRouter.GET("/getList", GetApiList)                 // 获取Api列表
		apiRouter.GET("/getAll", GetAllApis)                  // 获取所有api
		apiRouter.GET("/getApiById/:id", GetApiById)          // 获取单条Api消息
		apiRouter.POST("/createApi", CreateApi)               // 创建Api
		apiRouter.DELETE("/deleteApi/:id", DeleteApi)         // 删除Api
		apiRouter.PUT("/updateApi/:id", UpdateApi)            // 更新api
		apiRouter.DELETE("/deleteApisByIds", DeleteApisByIds) // 删除选中api
	}
}
