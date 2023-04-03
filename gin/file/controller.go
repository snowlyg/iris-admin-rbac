package file

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/iris-admin/server/web/web_gin/response"
	"github.com/snowlyg/iris-admin/server/zap_server"
)

func UploadFile(ctx *gin.Context) {
	_, header, err := ctx.Request.FormFile("file")
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	url, err := uploadFile(header) // 文件上传后拿到文件路径
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.OkWithDetailed(gin.H{"src": url}, "上传成功", ctx)
}
