package auth

import (
	"errors"
	"time"

	"github.com/snowlyg/iris-admin-rbac/iris/user"
	"github.com/snowlyg/iris-admin/server/database"
	"github.com/snowlyg/iris-admin/server/web"
	"github.com/snowlyg/iris-admin/server/zap_server"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrUserNameOrPassword = errors.New("用户名或密码错误")
)

// GetAccessToken 登录
func GetAccessToken(req *LoginRequest) (string, uint, error) {
	admin, err := user.FindPasswordByUserName(database.Instance(), req.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", 0, err
	}

	if admin == nil || admin.Id == 0 {
		return "", 0, ErrUserNameOrPassword
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password)); err != nil {
		zap_server.ZAPLOG.Error("用户名或密码错误", zap.String("密码:", req.Password), zap.String("hash:", admin.Password), zap.String("bcrypt.CompareHashAndPassword()", err.Error()))
		return "", 0, ErrUserNameOrPassword
	}
	expiresAt := time.Now().Local().Add(time.Duration(web.CONFIG.SessionTimeout) * time.Minute).Unix()
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
		return "", 0, err
	}

	return token, admin.Id, nil
}
