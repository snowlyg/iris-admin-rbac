package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/iris-admin/server/web/web_gin/middleware"
)

func Group(group *gin.RouterGroup) {
	adminRouter := group.Group("/admin", middleware.Auth(), middleware.CasbinHandler(), middleware.Cors())
	{
		adminRouter.GET("/getAll", GetAll)
		adminRouter.GET("/{id:uint}", GetAdmin)
		adminRouter.POST("/create", CreateAdmin)
		adminRouter.POST("/{id:uint}", UpdateAdmin)
		adminRouter.DELETE("/{id:uint}", DeleteAdmin)
		adminRouter.GET("/profile", Profile)
		adminRouter.POST("/change_avatar", ChangeAvatar)
	}
}
