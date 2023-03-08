package authority

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/iris-admin/server/database/orm"
	"github.com/snowlyg/iris-admin/server/zap_server"
)

type ReqPaginate struct {
	orm.Paginate
	AuthorityName string `json:"authorityName" form:"authorityName"`
	AuthorityType int    `json:"authorityType" form:"authorityType"`
}

type CreateAuthorityRequest struct {
	Uuid  string     `json:"uuid"`
	Perms [][]string `json:"perms" gorm:"-"`
	BaseAuthority
}

func (req *CreateAuthorityRequest) Request(ctx *gin.Context) error {
	if err := ctx.ShouldBindJSON(req); err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return orm.ErrParamValidate
	}
	return nil
}

type UpdateAuthorityRequest struct {
	BaseAuthority
}

func (req *UpdateAuthorityRequest) Request(ctx *gin.Context) error {
	if err := ctx.ShouldBindJSON(req); err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return orm.ErrParamValidate
	}
	return nil
}
