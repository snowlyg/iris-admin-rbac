package perm

import (
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/snowlyg/helper/str"
	"github.com/snowlyg/iris-admin/server/database"
	"github.com/snowlyg/iris-admin/server/database/orm"
	"github.com/snowlyg/iris-admin/server/database/scope"
	"github.com/snowlyg/iris-admin/server/web/web_gin/response"
)

// First 详情
func First(ctx iris.Context) {
	req := &orm.ReqId{}
	if err := req.Request(ctx); err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}
	perm := &Response{}
	err := orm.First(database.Instance(), perm, scope.IdScope(req.Id))
	if err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(orm.Response{Status: http.StatusOK, Data: perm, Msg: response.ResponseOkMessage})
}

// CreatePerm 添加
func CreatePerm(ctx iris.Context) {
	req := &Request{}
	if err := req.Request(ctx); err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}
	if !CheckNameAndAct(NameScope(req.Name), ActScope(req.Act)) {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: str.Join("权限[", req.Name, "-", req.Act, "]已存在")})
		return
	}
	perm := &Permission{BasePermission: req.BasePermission}
	id, err := orm.Create(database.Instance(), perm)
	if err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(orm.Response{Status: http.StatusOK, Data: iris.Map{"id": id}, Msg: response.ResponseOkMessage})
}

// UpdatePerm 更新
func UpdatePerm(ctx iris.Context) {
	reqId := &orm.ReqId{}
	if err := reqId.Request(ctx); err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}
	req := &Request{}
	if err := req.Request(ctx); err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}
	if !CheckNameAndAct(NameScope(req.Name), ActScope(req.Act), scope.NeIdScope(reqId.Id)) {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: str.Join("权限[", req.Name, "-", req.Act, "]已存在")})
		return
	}
	perm := &Permission{BasePermission: req.BasePermission}
	err := orm.Update(database.Instance(), reqId.Id, perm)
	if err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(orm.Response{Status: http.StatusOK, Data: nil, Msg: response.ResponseOkMessage})
}

// DeletePerm 删除
func DeletePerm(ctx iris.Context) {
	reqId := &orm.ReqId{}
	if err := reqId.Request(ctx); err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}
	err := orm.Delete(database.Instance(), &Permission{}, scope.IdScope(reqId.Id))
	if err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(orm.Response{Status: http.StatusOK, Data: nil, Msg: response.ResponseOkMessage})
}

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
