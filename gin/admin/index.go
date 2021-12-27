package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/iris-admin-rbac/gin/middleware"
)

func Group(group *gin.RouterGroup) {
	adminRouter := group.Group("/admin", middleware.Auth(), middleware.CasbinHandler())
	{
		adminRouter.GET("/getAll", GetAll)
		adminRouter.GET("/getAdmin/:id", GetAdmin)
		adminRouter.POST("/createAdmin", CreateAdmin)
		adminRouter.PUT("/updateAdmin/:id", UpdateAdmin)
		adminRouter.DELETE("/deleteAdmin/:id", DeleteAdmin)
		adminRouter.GET("/profile", Profile)
		adminRouter.POST("/changeAvatar", ChangeAvatar)
	}
}
