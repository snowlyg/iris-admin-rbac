package authority

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/iris-admin/server/database"
	"github.com/snowlyg/iris-admin/server/database/orm"
	"github.com/snowlyg/iris-admin/server/database/scope"
	"github.com/snowlyg/iris-admin/server/web/web_gin/response"
	"github.com/snowlyg/iris-admin/server/zap_server"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

// CreateAuthority 创建角色
func CreateAuthority(ctx *gin.Context) {
	req := &AuthorityRequest{}
	if errs := req.Request(ctx); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if id, err := Create(req); err != nil {
		zap_server.ZAPLOG.Error("添加角色数据失败", zap.Any("Create", err))
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.OkWithData(gin.H{"id": id}, ctx)
	}
}

// CopyAuthority 拷贝角色
func CopyAuthority(ctx *gin.Context) {
	reqId := &orm.ReqId{}
	if errs := ctx.ShouldBindUri(&reqId); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	req := &AuthorityRequest{}
	if errs := req.Request(ctx); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	if id, err := Copy(reqId.Id, req); err != nil {
		zap_server.ZAPLOG.Error("拷贝失败!", zap.Any("err", err))
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.OkWithData(gin.H{"id": id}, ctx)
	}
}

// UpdateAuthority 更新角色信息
func UpdateAuthority(ctx *gin.Context) {
	reqId := &orm.ReqId{}
	if errs := ctx.ShouldBindUri(&reqId); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	req := &AuthorityRequest{}
	if errs := req.Request(ctx); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	admin := &Authority{BaseAuthority: req.BaseAuthority}
	err := Update(reqId.Id, admin)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.Ok(ctx)
}

// GetAdminAuthorityList 分页获取管理角色列表
func GetAdminAuthorityList(ctx *gin.Context) {
	req := &orm.Paginate{}
	if errs := ctx.ShouldBind(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	items := &PageResponse{}
	total, err := orm.Pagination(database.Instance(), items, req.PaginateScope(), AuthorityTypeScope(multi.AdminAuthority))
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

// GetTenancyAuthorityList 分页获取商户角色列表
func GetTenancyAuthorityList(ctx *gin.Context) {
	req := &orm.Paginate{}
	if errs := ctx.ShouldBind(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	items := &PageResponse{}
	total, err := orm.Pagination(database.Instance(), items, req.PaginateScope(), AuthorityTypeScope(multi.TenancyAuthority))
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

// GetGeneralAuthorityList 分页获取用户角色列表
func GetGeneralAuthorityList(ctx *gin.Context) {
	req := &orm.Paginate{}
	if errs := ctx.ShouldBind(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	items := &PageResponse{}
	total, err := orm.Pagination(database.Instance(), items, req.PaginateScope(), AuthorityTypeScope(multi.GeneralAuthority))
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

// GetAuthorityList 分页获取所有角色列表
func GetAuthorityList(ctx *gin.Context) {
	req := &orm.Paginate{}
	if errs := ctx.ShouldBind(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	items := &PageResponse{}
	total, err := orm.Pagination(database.Instance(), items, req.PaginateScope())
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

// DeleteAuthority 删除角色
func DeleteAuthority(ctx *gin.Context) {
	reqId := &orm.ReqId{}
	if errs := ctx.ShouldBindUri(&reqId); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	err := orm.Delete(database.Instance(), &Authority{}, scope.IdScope(reqId.Id))
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.Ok(ctx)
}
