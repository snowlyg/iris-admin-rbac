package admin

import (
	"errors"
	"strconv"

	"github.com/snowlyg/helper/arr"
	"github.com/snowlyg/iris-admin-rbac/gin/authority"
	"github.com/snowlyg/iris-admin/server/casbin"
	"github.com/snowlyg/iris-admin/server/database"
	"github.com/snowlyg/iris-admin/server/database/orm"
	"github.com/snowlyg/iris-admin/server/database/scope"
	"github.com/snowlyg/iris-admin/server/zap_server"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var ErrUserNameInvalide = errors.New("用户名名称已经被使用")

// transform
func transform(admins ...*Response) {
	var roleUuids []string
	userRoleUuids := map[uint][]string{}
	if len(admins) == 0 {
		return
	}
	for _, admin := range admins {
		admin.ToString()
		userRUuids := casbin.GetRolesForUser(admin.Id)
		userRoleUuids[admin.Id] = userRUuids
		roleUuids = append(roleUuids, userRUuids...)
	}

	roles, err := authority.FindInUuid(database.Instance(), roleUuids)
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return
	}
	if len(roles) == 0 {
		return
	}
	for _, admin := range admins {
		for _, role := range roles {
			if arr.InArray(userRoleUuids[admin.Id], role.Uuid) {
				admin.Authorities = append(admin.Authorities, role.AuthorityName)
			}
		}
	}
}

// FindByName
func FindByUserName(scopes ...func(db *gorm.DB) *gorm.DB) (*Response, error) {
	admin := &Response{}
	err := admin.First(database.Instance(), scopes...)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

func FindPasswordByUserName(db *gorm.DB, username string, ids ...uint) (*LoginResponse, error) {
	admin := &LoginResponse{}
	db = db.Model(&Admin{}).Select("id,password").
		Where("username = ?", username)
	if len(ids) == 1 {
		db.Where("id != ?", ids[0])
	}
	err := db.First(admin).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return admin, err
	}
	userId := strconv.FormatUint(uint64(admin.Id), 10)
	admin.AuthorityIds, err = getUserRoleUuids(userId)
	if err != nil {
		return nil, err
	}
	return admin, nil
}

// getUserRoleUuids
func getUserRoleUuids(userId string) ([]string, error) {
	roleUuids, err := casbin.Instance().GetRolesForUser(userId)
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return nil, err
	}
	return roleUuids, nil
}

func Create(req *Request) (uint, error) {
	if _, err := FindByUserName(UserNameScope(req.Username)); !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, ErrUserNameInvalide
	}
	admin := &Admin{BaseAdmin: req.BaseAdmin, AuthorityIds: req.AuthorityUuids}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return 0, err
	}

	zap_server.ZAPLOG.Info("添加用户", zap.String("hash:", req.Password), zap.ByteString("hash:", hash))

	admin.Password = string(hash)
	id, err := orm.Create(database.Instance(), admin)
	if err != nil {
		return 0, err
	}

	if err := AddRoleForUser(admin); err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return 0, err
	}

	return id, nil
}

func IsAdminUser(id uint) error {
	admin := &Response{}
	err := orm.First(database.Instance(), admin, scope.IdScope(id))
	if err != nil {
		return err
	}
	if arr.InArray(admin.Authorities, authority.GetAdminRoleName()) {
		return errors.New("不能操作超级管理员")
	}
	return nil
}

// AddRoleForUser add roles for user
func AddRoleForUser(admin *Admin) error {
	userId := strconv.FormatUint(uint64(admin.ID), 10)
	oldRoleUuids, err := getUserRoleUuids(userId)
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return err
	}

	if len(oldRoleUuids) > 0 {
		if _, err := casbin.Instance().DeleteRolesForUser(userId); err != nil {
			zap_server.ZAPLOG.Error(err.Error())
			return err
		}
	}
	if len(admin.AuthorityIds) == 0 {
		return nil
	}

	var roleUuids []string
	roleUuids = append(roleUuids, admin.AuthorityIds...)

	if _, err := casbin.Instance().AddRolesForUser(userId, roleUuids); err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return err
	}

	return nil
}

func UpdateAvatar(db *gorm.DB, id uint, avatar string) error {
	err := db.Model(&Admin{}).Where("id = ?", id).Update("header_img", avatar).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return err
	}

	return nil
}
