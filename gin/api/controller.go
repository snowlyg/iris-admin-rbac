package api

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/iris-admin/server/database"
	"github.com/snowlyg/iris-admin/server/database/scope"
	"github.com/snowlyg/iris-admin/server/web/web_gin/request"
	"github.com/snowlyg/iris-admin/server/web/web_gin/response"
	"gorm.io/gorm"
)

// CreateApi 创建基础api
func CreateApi(ctx *gin.Context) {
	api := &Api{}
	if errs := ctx.ShouldBindJSON(api); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if id, err := api.Create(database.Instance()); err != nil {
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.OkWithData(gin.H{"id": id, "path": api.Path, "method": api.Method}, ctx)
	}
}

// DeleteApi 删除api
func DeleteApi(ctx *gin.Context) {
	var reqId request.IdBinding
	if errs := ctx.ShouldBindUri(&reqId); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	var req DeleteApiReq
	if errs := ctx.ShouldBind(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := Delete(reqId.Id, req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.Ok(ctx)
	}
}

// GetApiList 分页获取API列表
func GetApiList(ctx *gin.Context) {
	var pageInfo ReqPaginate
	if errs := ctx.ShouldBind(&pageInfo); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	items := &PageResponse{}
	scopes := []func(db *gorm.DB) *gorm.DB{IsApiScope()}
	if pageInfo.AuthorityType.AuthorityType > 0 {
		scopes = append(scopes, AuthorityTypeScope(pageInfo.AuthorityType.AuthorityType))
	}
	if pageInfo.Method != "" {
		scopes = append(scopes, MethodScope(pageInfo.Method))
	}
	if pageInfo.ApiGroup != "" {
		scopes = append(scopes, ApiGroupScope(pageInfo.ApiGroup))
	}
	if pageInfo.Path != "" {
		scopes = append(scopes, PathScope(pageInfo.Path))
	}
	total, err := items.Paginate(database.Instance(), pageInfo.PaginateScope(), scopes...)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.OkWithData(response.PageResult{
			List:     items.Item,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, ctx)
	}
}

// GetApiById 根据id获取api
func GetApiById(ctx *gin.Context) {
	var req request.IdBinding
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	api := &Response{}
	err := api.First(database.Instance(), scope.IdScope(req.Id))
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.OkWithData(api, ctx)
	}
}

// UpdateApi 更新基础api
func UpdateApi(ctx *gin.Context) {
	var req request.IdBinding
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	api := &Api{}
	if errs := ctx.ShouldBindJSON(api); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	err := api.Update(database.Instance(), scope.IdScope(req.Id))
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.Ok(ctx)
	}
}

// GetAllApis 获取所有的Api不分页
func GetAllApis(ctx *gin.Context) {
	var req AuthorityType
	if err := ctx.ShouldBind(&req); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	res := &PageResponse{}
	err := res.Find(database.Instance(), AuthorityTypeScope(req.AuthorityType), IsApiScope())
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	routers := FormatApis(res.Item)

	response.OkWithData(routers, ctx)

}

// DeleteApisByIds 删除选中Api
func DeleteApisByIds(ctx *gin.Context) {
	var reqIds request.IdsBinding
	if errs := ctx.ShouldBindJSON(&reqIds); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := BatcheDelete(reqIds.Ids); err != nil {
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.Ok(ctx)
	}
}
