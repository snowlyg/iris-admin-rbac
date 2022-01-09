package file

import (
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/snowlyg/iris-admin/server/database/orm"
	"github.com/snowlyg/iris-admin/server/web/web_gin/response"
	"github.com/snowlyg/iris-admin/server/zap_server"
	"go.uber.org/zap"
)

// Upload 上传文件
// - 需要 file 表单文件字段
func Upload(ctx iris.Context) {
	f, fh, err := ctx.FormFile("file")
	if err != nil {
		zap_server.ZAPLOG.Error("文件上传失败", zap.String("ctx.FormFile(\"file\")", err.Error()))
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}
	defer f.Close()

	data, err := UploadFile(ctx, fh)
	if err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(orm.Response{Status: http.StatusOK, Data: data, Msg: response.ResponseOkMessage})
}
