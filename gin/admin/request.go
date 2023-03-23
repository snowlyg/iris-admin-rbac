package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/iris-admin/server/database/orm"
	"github.com/snowlyg/iris-admin/server/zap_server"
)

type Request struct {
	BaseAdmin
	Password       string   `json:"password"`
	AuthorityUuids []string `json:"authorityIds" form:"authorityIds"`
}

func (req *Request) Request(ctx *gin.Context) error {
	if err := ctx.ShouldBindJSON(req); err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return orm.ErrParamValidate
	}
	return nil
}

type ReqPaginate struct {
	orm.Paginate
	SearchKey string `json:"searchKey" form:"searchKey" uri:"searchKey" param:"searchKey"`
	Status    int    `json:"status" form:"status" uri:"status" param:"status"`
}
