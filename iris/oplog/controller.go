package oplog

import (
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/snowlyg/iris-admin/server/database"
	"github.com/snowlyg/iris-admin/server/database/orm"
	"github.com/snowlyg/iris-admin/server/web/web_gin/response"
)

// GetAll 分页列表
// - 获取分页参数
// - 请求分页数据
func GetAll(ctx iris.Context) {
	req := &orm.Paginate{}
	if err := req.Request(ctx); err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}
	items := &PageResponse{}
	total, err := orm.Pagination(database.Instance(), items, req.PaginateScope())
	if err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}
	list := iris.Map{"items": items.Item, "total": total, "pageSize": req.PageSize, "page": req.Page}
	ctx.JSON(orm.Response{Status: http.StatusOK, Data: list, Msg: response.ResponseOkMessage})
}
