package public

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/iris-admin/server/database/orm"
	"github.com/snowlyg/iris-admin/server/zap_server"
)

// LoginRequest 登录
type LoginRequest struct {
	Username      string `json:"username" form:"username" binding:"required"`
	Password      string `json:"password" form:"password" binding:"required"`
	Captcha       string `json:"captcha" form:"captcha" binding:"dev-required"`
	CaptchaId     string `json:"captchaId" form:"captchaId" binding:"dev-required"`
	AuthorityType int    `json:"authorityType" `
}

func (req *LoginRequest) Request(ctx *gin.Context) error {
	if err := ctx.ShouldBindJSON(req); err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return orm.ErrParamValidate
	}
	return nil
}