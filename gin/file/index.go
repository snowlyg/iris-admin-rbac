package file

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/iris-admin-rbac/gin/middleware"
)

func Group(group *gin.RouterGroup) {
	apiRouter := group.Group("/file", middleware.Auth(), middleware.CasbinHandler())
	{
		apiRouter.POST("/upload", UploadFile)
	}
}
