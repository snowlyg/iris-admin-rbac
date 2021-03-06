package auth

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/iris-admin/server/database/orm"
	"github.com/snowlyg/iris-admin/server/zap_server"
)

// LoginRequest 登录请求字段
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (req *LoginRequest) Request(ctx iris.Context) error {
	if err := ctx.ReadJSON(req); err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return orm.ErrParamValidate
	}
	return nil
}
