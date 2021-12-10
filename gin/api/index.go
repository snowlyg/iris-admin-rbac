package api

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/iris-admin/server/web/web_gin/middleware"
)

func Group(group *gin.RouterGroup) {
	apiRouter := group.Group("/api", middleware.Auth(), middleware.CasbinHandler(), middleware.Cors())
	{
		apiRouter.GET("/getApiList", GetApiList)              // 获取Api列表
		apiRouter.GET("/getAllApis", GetAllApis)              // 获取所有api
		apiRouter.GET("/getApiById/:id", GetApiById)          // 获取单条Api消息
		apiRouter.POST("/createApi", CreateApi)               // 创建Api
		apiRouter.DELETE("/deleteApi", DeleteApi)             // 删除Api
		apiRouter.PUT("/updateApi/:id", UpdateApi)            // 更新api
		apiRouter.DELETE("/deleteApisByIds", DeleteApisByIds) // 删除选中api
	}
}
