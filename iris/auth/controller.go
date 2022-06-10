package auth

import (
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/snowlyg/iris-admin/server/database/orm"
	"github.com/snowlyg/iris-admin/server/web/web_gin/response"
)

// Login 登录
// - LoginRequest 登录需要提交 uesrname 和 password 字段到接口
// - validate.ValidRequest 验证接口提交参数，需要在 LoginRequest 的字段设置 validate:"required"
// - GetAccessToken 生成验证 token
func Login(ctx iris.Context) {
	req := &LoginRequest{}
	if err := req.Request(ctx); err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}
	token, id, err := GetAccessToken(req)
	if err != nil {
		ctx.JSON(orm.Response{Status: http.StatusBadRequest, Data: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(orm.Response{Status: http.StatusOK, Data: iris.Map{"accessToken": token, "user": iris.Map{"id": id}}, Msg: response.ResponseOkMessage})
}
