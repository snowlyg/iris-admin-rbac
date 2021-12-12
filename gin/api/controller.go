package api

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/iris-admin/server/database"
	"github.com/snowlyg/iris-admin/server/database/orm"
	"github.com/snowlyg/iris-admin/server/database/scope"
	"github.com/snowlyg/iris-admin/server/web/web_gin/request"
	"github.com/snowlyg/iris-admin/server/web/web_gin/response"
	"github.com/snowlyg/iris-admin/server/zap_server"
	"go.uber.org/zap"
)

// CreateApi 创建基础api
func CreateApi(ctx *gin.Context) {
	api := &Api{}
	if errs := ctx.ShouldBindJSON(api); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if id, err := orm.Create(database.Instance(), api); err != nil {
		zap_server.ZAPLOG.Error("添加角色数据失败", zap.Any("Create", err))
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.OkWithData(gin.H{"id": id, "path": api.Path, "method": api.Method}, ctx)
	}
}

// DeleteApi 删除api
func DeleteApi(ctx *gin.Context) {
	var reqId request.GetById
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
		zap_server.ZAPLOG.Error("删除失败", zap.Any("Delete", err))
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
	total, err := orm.Pagination(database.Instance(), items, pageInfo.PaginateScope())
	if err != nil {
		zap_server.ZAPLOG.Error("获取分页获取API列表失败", zap.Any("orm.Pagination", err))
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
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	api := &Response{}
	err := api.First(database.Instance(), scope.IdScope(req.Id))
	if err != nil {
		zap_server.ZAPLOG.Error("查询权限失败", zap.Any("api.First", err))
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.OkWithData(api, ctx)
	}
}

// UpdateApi 更新基础api
func UpdateApi(ctx *gin.Context) {
	var req request.GetById
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
		zap_server.ZAPLOG.Error("修改失败", zap.Any("api.Update", err))
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
	apis := &PageResponse{}
	err := apis.Find(database.Instance(), AuthorityTypeScope(req.AuthorityType))
	if err != nil {
		zap_server.ZAPLOG.Error("获取所有的Api不分页", zap.Any("orm.Find", err))
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.OkWithData(apis.Item, ctx)
	}
}

// DeleteApisByIds 删除选中Api
func DeleteApisByIds(ctx *gin.Context) {
	var reqIds request.Ids
	if errs := ctx.ShouldBindJSON(&reqIds); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := BatcheDelete(reqIds.Ids); err != nil {
		zap_server.ZAPLOG.Error("批量删除", zap.Any("BatcheDelete", err))
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.Ok(ctx)
	}
}
