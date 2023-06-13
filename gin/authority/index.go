package authority

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/iris-admin-rbac/gin/middleware"
)

func Group(group *gin.RouterGroup) {

	authRouter := group.Group("/authority", middleware.Auth(), middleware.CasbinHandler())
	{
		authRouter.GET("/getAuthorityList", GetAuthorityList)               // 获取角色列表
		authRouter.GET("/getAdminAuthorityList", GetAdminAuthorityList)     // 获取员工角色列表
		authRouter.GET("/getTenancyAuthorityList", GetTenancyAuthorityList) // 获取商户角色列表
		authRouter.GET("/getGeneralAuthorityList", GetGeneralAuthorityList) // 获取普通用户角色列表
		authRouter.POST("/createAuthority", CreateAuthority)                // 创建角色
		authRouter.PUT("/updateAuthority/:id", UpdateAuthority)             // 更新角色
		authRouter.POST("/copyAuthority/:id", CopyAuthority)                // 复制
		authRouter.DELETE("/deleteAuthority/:id", DeleteAuthority)          // 删除角色
		authRouter.GET("/authorityDetail/:id", AuthorityDetail)             // 角色权限
	}
}
