package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/iris-admin-rbac/gin/admin"
	"github.com/snowlyg/iris-admin-rbac/gin/api"
	"github.com/snowlyg/iris-admin-rbac/gin/authority"
	"github.com/snowlyg/iris-admin-rbac/gin/public"
)

// Party v1 模块
func Party(group *gin.RouterGroup) {
	api.Group(group)
	admin.Group(group)
	authority.Group(group)
	public.Group(group)
}
