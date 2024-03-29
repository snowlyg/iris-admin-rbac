package public

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/snowlyg/iris-admin/server/web"
	"github.com/snowlyg/iris-admin/server/web/web_gin/response"
	"github.com/snowlyg/multi"
	multi_gin "github.com/snowlyg/multi/gin"
)

var store = base64Captcha.DefaultMemStore

// AdminLogin 后台登录
func AdminLogin(ctx *gin.Context) {
	req := &LoginRequest{}
	if errs := req.Request(ctx); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	req.AuthorityType = multi.AdminAuthority
	token, err := GetAccessToken(req)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.OkWithData(token, ctx)
	}
}

// Logout 退出
func Logout(ctx *gin.Context) {
	token := multi_gin.GetVerifiedToken(ctx)
	if token == nil {
		response.FailWithMessage("授权凭证为空", ctx)
		return
	}
	err := DelToken(string(token))
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.Ok(ctx)
}

// Clear 清空 token
func Clear(ctx *gin.Context) {
	token := multi_gin.GetVerifiedToken(ctx)
	if token == nil {
		response.FailWithMessage("授权凭证为空", ctx)
		return
	}
	if err := CleanToken(multi.AdminAuthority, string(token)); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.Ok(ctx)
}

// Captcha 生成验证码
func Captcha(ctx *gin.Context) {
	//字符,公式,验证码配置
	// 生成默认数字的driver
	if web.CONFIG.System.Level != "release" && web.CONFIG.Captcha.KeyLong <= 0 {
		response.OkWithData(gin.H{"enable": false}, ctx)
		return
	}
	keyLong := web.CONFIG.Captcha.KeyLong
	if web.CONFIG.System.Level == "release" {
		keyLong = 4
	}

	driver := base64Captcha.NewDriverDigit(web.CONFIG.Captcha.ImgHeight, web.CONFIG.Captcha.ImgWidth, keyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		response.FailWithMessage(err.Error(), ctx)
	} else {
		response.OkWithData(gin.H{"captchaId": id, "picPath": b64s, "enable": true}, ctx)
	}
}
