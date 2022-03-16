package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/iris-admin/server/database/orm"
	"github.com/snowlyg/iris-admin/server/zap_server"
)

type Request struct {
	BaseAdmin
	Password     string   `json:"password"`
	AuthorityUuids []string `json:"authorityIds"`
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
	Name string `json:"name"`
}
