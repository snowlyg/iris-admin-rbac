package role

import (
	"errors"
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/snowlyg/iris-admin/server/database"
	"github.com/snowlyg/iris-admin/server/database/orm"
	"github.com/snowlyg/iris-admin/server/database/scope"
	"github.com/snowlyg/iris-admin/server/web/web_gin/response"
	"gorm.io/gorm"
)

// First 详情
func First(ctx iris.Context) {
	req := &orm.ReqId{}
	if err := req.Request(ctx); err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}
	perm := &Response{}
	err := perm.First(database.Instance(), scope.IdScope(req.Id))
	if err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(orm.Response{Status: http.StatusOK, Data: perm, Msg: response.ResponseOkMessage})
}

// CreateRole 添加
func CreateRole(ctx iris.Context) {
	req := &Request{}
	if err := req.Request(ctx); err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}

	id, err := Create(req)
	if err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
	}

	ctx.JSON(orm.Response{Status: http.StatusOK, Data: iris.Map{"id": id}, Msg: response.ResponseOkMessage})
}

// UpdateRole 更新
func UpdateRole(ctx iris.Context) {
	reqId := &orm.ReqId{}
	if err := reqId.Request(ctx); err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}

	err := IsAdminRole(reqId.Id)
	if err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
	}

	req := &Request{}
	if err := req.Request(ctx); err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}

	if _, err := FindByName(NameScope(req.Name), scope.NeIdScope(reqId.Id)); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: "角色名称已经被使用"})
		return
	}

	role := &Role{BaseRole: req.BaseRole}
	err = role.Update(database.Instance(), scope.IdScope(reqId.Id))
	if err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}

	err = AddPermForRole(req.Name, req.Perms)
	if err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
	}

	ctx.JSON(orm.Response{Status: http.StatusOK, Data: nil, Msg: response.ResponseOkMessage})
}

// DeleteRole 删除
func DeleteRole(ctx iris.Context) {
	reqId := &orm.ReqId{}
	if err := reqId.Request(ctx); err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}

	err := IsAdminRole(reqId.Id)
	if err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
	}
	role := &Role{}
	err = role.Delete(database.Instance(), scope.IdScope(reqId.Id))
	if err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}

	ctx.JSON(orm.Response{Status: http.StatusOK, Data: nil, Msg: response.ResponseOkMessage})
}

// GetAll 分页列表
func GetAll(ctx iris.Context) {
	req := &orm.Paginate{}
	if err := req.Request(ctx); err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}

	items := &PageResponse{}
	total, err := items.Paginate(database.Instance(), req.PaginateScope())
	if err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}
	list := iris.Map{"items": items.Item, "total": total, "pageSize": req.PageSize, "page": req.Page}
	ctx.JSON(orm.Response{Status: http.StatusOK, Data: list, Msg: response.ResponseOkMessage})
}
