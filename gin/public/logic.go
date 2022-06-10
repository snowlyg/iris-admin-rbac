package public

import (
	"errors"
	"time"

	"github.com/snowlyg/iris-admin-rbac/gin/admin"
	"github.com/snowlyg/iris-admin/server/database"
	"github.com/snowlyg/iris-admin/server/web"
	"github.com/snowlyg/iris-admin/server/zap_server"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNameOrPassword = errors.New("用户名或密码错误")
	ErrCaptcha            = errors.New("验证码错误")
)

// GetAccessToken 登录
func GetAccessToken(req *LoginRequest) (*LoginResponse, error) {
	if !store.Verify(req.CaptchaId, req.Captcha, true) && web.CONFIG.System.Level != "test" {
		return nil, ErrCaptcha
	}
	admin, err := admin.FindPasswordByUserName(database.Instance(), req.Username)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password)); err != nil {
		zap_server.ZAPLOG.Error("用户名或密码错误:",
			zap.String("密码:", req.Password),
			zap.String("hash:", admin.Password),
			zap.String("错误:", err.Error()))
		return nil, ErrUserNameOrPassword
	}
	expiresAt := time.Now().Local().Add(multi.RedisSessionTimeoutWeb).Unix()
	claims := multi.New(&multi.Multi{
		Id:            admin.Id,
		Username:      req.Username,
		AuthorityIds:  admin.AuthorityIds,
		AuthorityType: multi.AdminAuthority,
		LoginType:     multi.LoginTypeWeb,
		AuthType:      multi.AuthPwd,
		ExpiresAt:     expiresAt,
	})
	token, _, err := multi.AuthDriver.GenerateToken(claims)
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return nil, err
	}
	if token == "" {
		zap_server.ZAPLOG.Error(multi.ErrEmptyToken.Error())
		return nil, multi.ErrEmptyToken
	}
	loginResponse := &LoginResponse{
		User: map[string]interface{}{
			"id": admin.Id,
		},
		Token: token,
	}
	return loginResponse, nil
}

// DelToken 删除token
func DelToken(token string) error {
	err := multi.AuthDriver.DelUserTokenCache(token)
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return err
	}
	return nil
}

// CleanToken 清空 token
func CleanToken(authorityType int, userId string) error {
	err := multi.AuthDriver.CleanUserTokenCache(authorityType, userId)
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return err
	}
	return nil
}
