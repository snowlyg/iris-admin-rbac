package admin

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/iris-admin/server/database"
	"github.com/snowlyg/iris-admin/server/database/orm"
	"github.com/snowlyg/iris-admin/server/database/scope"
	"github.com/snowlyg/iris-admin/server/web/web_gin/response"
	multi "github.com/snowlyg/multi/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Profile 个人信息
func Profile(ctx *gin.Context) {
	item := &Response{}
	err := item.First(database.Instance(), scope.IdScope(multi.GetUserId(ctx)))
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithData(item, ctx)
}

// GetAdmin 详情
func GetAdmin(ctx *gin.Context) {
	req := &orm.ReqId{}
	if errs := req.Request(ctx); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	admin := &Response{}
	err := admin.First(database.Instance(), scope.IdScope(req.Id))
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithData(admin, ctx)
}

// CreateAdmin添加
func CreateAdmin(ctx *gin.Context) {
	req := &Request{}
	if errs := ctx.ShouldBindJSON(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	id, err := Create(req)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
	}
	response.OkWithData(gin.H{"id": id}, ctx)
}

// UpdateAdmin 更新
func UpdateAdmin(ctx *gin.Context) {
	reqId := &orm.ReqId{}
	if errs := reqId.Request(ctx); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if err := IsAdminUser(reqId.Id); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	req := &Request{}
	if errs := ctx.ShouldBindJSON(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if _, err := FindByUserName(UserNameScope(req.Username), scope.NeIdScope(reqId.Id)); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	admin := &Admin{BaseAdmin: req.BaseAdmin, AuthorityIds: req.AuthorityUuids}
	if req.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			response.FailWithMessage(err.Error(), ctx)
			return
		}
		admin.Password = string(hash)
	}
	err := admin.Update(database.Instance(), scope.IdScope(reqId.Id))
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	admin.ID = reqId.Id
	if err := AddRoleForUser(admin); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	response.Ok(ctx)
}

// DeleteAdmin 删除
func DeleteAdmin(ctx *gin.Context) {
	reqId := &orm.ReqId{}
	if errs := ctx.ShouldBindUri(&reqId); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	admin := &Admin{}
	err := admin.Delete(database.Instance(), scope.IdScope(reqId.Id))
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.Ok(ctx)
}

// GetAll 分页列表
func GetAll(ctx *gin.Context) {
	req := &ReqPaginate{}
	if errs := ctx.ShouldBind(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	scopes := []func(db *gorm.DB) *gorm.DB{}
	if req.SearchKey != "" {
		scopes = append(scopes, SearchKeyScope(req.SearchKey))
	}
	if req.Status != "" {
		scopes = append(scopes, StatusScope(req.Status))
	}

	items := &PageResponse{}
	total, err := items.Paginate(database.Instance(), req.PaginateScope(), scopes...)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithData(response.PageResult{
		List:     items.Item,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, ctx)
}

// ChangeAvatar 修改头像
func ChangeAvatar(ctx *gin.Context) {
	avatar := &Avatar{}
	if errs := ctx.ShouldBindJSON(&avatar); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	err := UpdateAvatar(database.Instance(), multi.GetUserId(ctx), avatar.HeaderImg)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.Ok(ctx)
}
