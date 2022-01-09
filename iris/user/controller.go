package user

import (
	"errors"
	"net/http"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/snowlyg/iris-admin/server/database"
	"github.com/snowlyg/iris-admin/server/database/orm"
	"github.com/snowlyg/iris-admin/server/database/scope"
	"github.com/snowlyg/iris-admin/server/web/web_gin/response"
	"github.com/snowlyg/iris-admin/server/web/web_iris/validate"
	"github.com/snowlyg/multi"
	multi_iris "github.com/snowlyg/multi/iris"
	"gorm.io/gorm"
)

// Profile 个人信息
func Profile(ctx iris.Context) {
	item := &Response{}
	err := orm.First(database.Instance(), item, scope.IdScope(multi_iris.GetUserId(ctx)))
	if err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(orm.Response{Status: http.StatusOK, Data: item, Msg: response.ResponseOkMessage})
}

// GetUser 详情
func GetUser(ctx iris.Context) {
	req := &orm.ReqId{}
	if err := req.Request(ctx); err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}
	user := &Response{}
	err := orm.First(database.Instance(), user, scope.IdScope(req.Id))
	if err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(orm.Response{Status: http.StatusOK, Data: user, Msg: response.ResponseOkMessage})
}

// CreateUser 添加
func CreateUser(ctx iris.Context) {
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

// UpdateUser 更新
func UpdateUser(ctx iris.Context) {
	reqId := &orm.ReqId{}
	if err := reqId.Request(ctx); err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}

	if err := IsAdminUser(reqId.Id); err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
	}

	req := &Request{}
	if err := req.Request(ctx); err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}

	if _, err := FindByUserName(UserNameScope(req.Username), scope.NeIdScope(reqId.Id)); !errors.Is(err, gorm.ErrRecordNotFound) {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
	}

	user := &User{BaseUser: req.BaseUser}
	err := orm.Update(database.Instance(), reqId.Id, user)
	if err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}

	if err := AddRoleForUser(user); err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
	}

	ctx.JSON(orm.Response{Status: http.StatusOK, Data: nil, Msg: response.ResponseOkMessage})
}

// DeleteUser 删除
func DeleteUser(ctx iris.Context) {
	reqId := &orm.ReqId{}
	if err := reqId.Request(ctx); err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}
	err := orm.Delete(database.Instance(), &User{}, scope.IdScope(reqId.Id))
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
	total, err := orm.Pagination(database.Instance(), items, req.PaginateScope())
	if err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}
	list := iris.Map{"items": items.Item, "total": total, "pageSize": req.PageSize, "page": req.Page}
	ctx.JSON(orm.Response{Status: http.StatusOK, Data: list, Msg: response.ResponseOkMessage})
}

// Logout 退出
func Logout(ctx iris.Context) {
	token := multi_iris.GetVerifiedToken(ctx)
	if token == nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: "授权凭证为空"})
		return
	}
	err := DelToken(string(token))
	if err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(orm.Response{Status: http.StatusOK, Data: nil, Msg: response.ResponseOkMessage})
}

// Clear 清空 token
func Clear(ctx iris.Context) {
	token := multi_iris.GetVerifiedToken(ctx)
	if token == nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: "授权凭证为空"})
		return
	}
	if err := CleanToken(multi.AdminAuthority, string(token)); err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(orm.Response{Status: http.StatusOK, Data: nil, Msg: response.ResponseOkMessage})
}

// ChangeAvatar 修改头像
func ChangeAvatar(ctx iris.Context) {
	avatar := &Avatar{}
	if err := ctx.ReadJSON(avatar); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: strings.Join(errs, ";")})
			return
		}
	}
	err := UpdateAvatar(database.Instance(), multi_iris.GetUserId(ctx), avatar.Avatar)
	if err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(orm.Response{Status: http.StatusOK, Data: nil, Msg: response.ResponseOkMessage})
}
