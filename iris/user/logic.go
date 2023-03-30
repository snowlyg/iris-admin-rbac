package user

import (
	"errors"
	"strconv"

	"github.com/snowlyg/helper/arr"
	"github.com/snowlyg/iris-admin-rbac/gin/authority"
	"github.com/snowlyg/iris-admin-rbac/iris/role"
	"github.com/snowlyg/iris-admin/server/casbin"
	"github.com/snowlyg/iris-admin/server/database"
	"github.com/snowlyg/iris-admin/server/database/scope"
	"github.com/snowlyg/iris-admin/server/zap_server"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var ErrUserNameInvalide = errors.New("用户名名称已经被使用")

// getRoles
func getRoles(db *gorm.DB, users ...*Response) {
	var roleNames []string
	userRoleNames := map[uint][]string{}
	if len(users) == 0 {
		return
	}
	for _, user := range users {
		user.ToString()
		userRoleName := casbin.GetRolesForUser(user.Id)
		userRoleNames[user.Id] = userRoleName
		roleNames = append(roleNames, userRoleName...)
	}

	roles, _ := role.FindInName(db, roleNames)
	if len(roles) > 0 {
		for _, user := range users {
			for _, role := range roles {
				userRoleNameLen := len(userRoleNames[user.Id])
				if userRoleNameLen > 0 {
					roleUuidType := arr.NewCheckArrayType(userRoleNameLen)
					for _, v := range userRoleNames[user.Id] {
						roleUuidType.Add(v)
					}
					if roleUuidType.Check(role.Name) {
						user.Roles = append(user.Roles, role.DisplayName)
					}
				}
			}
		}
	}
}

// FindByName
func FindByUserName(scopes ...func(db *gorm.DB) *gorm.DB) (*Response, error) {
	user := &Response{}
	err := user.First(database.Instance(), scopes...)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func FindPasswordByUserName(db *gorm.DB, username string, scopes ...func(db *gorm.DB) *gorm.DB) (*LoginResponse, error) {
	user := &LoginResponse{}
	if db == nil {
		return nil, gorm.ErrInvalidDB
	}
	db = db.Model(&User{}).Select("id,password").
		Where("username = ?", username)

	if len(scopes) > 0 {
		db.Scopes(scopes...)
	}

	err := db.First(user).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return nil, err
	}
	userId := strconv.FormatUint(uint64(user.Id), 10)
	user.AuthorityIds, err = getUserRoleNames(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// getUserRoleNames
func getUserRoleNames(userId string) ([]string, error) {
	roleNames, err := casbin.Instance().GetRolesForUser(userId)
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return nil, err
	}
	return roleNames, nil
}

func Create(req *Request) (uint, error) {
	if _, err := FindByUserName(UserNameScope(req.Username)); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, ErrUserNameInvalide
	}
	user := &User{BaseUser: req.BaseUser, RoleNames: req.RoleNames}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return 0, err
	}

	zap_server.ZAPLOG.Info("添加用户", zap.String("hash:", req.Password), zap.ByteString("hash:", hash))

	user.Password = string(hash)
	id, err := user.Create(database.Instance())
	if err != nil {
		return 0, err
	}

	if err := AddRoleForUser(user); err != nil {
		return 0, err
	}

	return id, nil
}

func IsAdminUser(id uint) error {
	user := &Response{}
	err := user.First(database.Instance(), scope.IdScope(id))
	if err != nil {
		return err
	}
	roleType := arr.NewCheckArrayType(len(user.Roles))
	for _, userRole := range user.Roles {
		roleType.Add(userRole)
	}
	if roleType.Check(authority.GetAdminRoleName()) {
		return errors.New("不能操作超级管理员")
	}
	return nil
}

func FindById(db *gorm.DB, id uint) (Response, error) {
	user := Response{}
	if db == nil {
		return user, gorm.ErrInvalidDB
	}
	err := db.Model(&User{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return user, err
	}

	getRoles(db, &user)

	return user, nil
}

// AddRoleForUser add roles for user
func AddRoleForUser(user *User) error {
	userId := strconv.FormatUint(uint64(user.ID), 10)
	oldRoleNames, err := getUserRoleNames(userId)
	if err != nil {
		return err
	}

	if len(oldRoleNames) > 0 {
		if _, err := casbin.Instance().DeleteRolesForUser(userId); err != nil {
			zap_server.ZAPLOG.Error(err.Error())
			return err
		}
	}
	if len(user.RoleNames) == 0 {
		return nil
	}

	var roleNames []string
	roleNames = append(roleNames, user.RoleNames...)

	if _, err := casbin.Instance().AddRolesForUser(userId, roleNames); err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return err
	}

	return nil
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

func UpdateAvatar(db *gorm.DB, id uint, avatar string) error {
	if db == nil {
		return gorm.ErrInvalidDB
	}
	err := db.Model(&User{}).Where("id = ?", id).Update("avatar", avatar).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return err
	}

	return nil
}
