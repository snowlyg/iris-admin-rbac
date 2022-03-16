package role

import (
	"errors"
	"fmt"

	"github.com/snowlyg/iris-admin/server/casbin"
	"github.com/snowlyg/iris-admin/server/database"
	"github.com/snowlyg/iris-admin/server/database/orm"
	"github.com/snowlyg/iris-admin/server/database/scope"
	"github.com/snowlyg/iris-admin/server/zap_server"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var ErrRoleNameInvalide = errors.New("角色名称已经被使用")

// GetAdminRoleName 获管理员角色名称
func GetAdminRoleName() string {
	return "admin"
}

// Create 添加
func Create(req *Request) (uint, error) {
	if _, err := FindByName(NameScope(req.Name)); !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, ErrRoleNameInvalide
	}
	role := &Role{BaseRole: req.BaseRole}
	id, err := orm.Create(database.Instance(), role)
	if err != nil {
		return 0, err
	}
	err = AddPermForRole(req.Name, req.Perms)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// FindByName
func FindByName(scopes ...func(db *gorm.DB) *gorm.DB) (*Response, error) {
	role := &Response{}
	err := orm.First(database.Instance(), role, scopes...)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func IsAdminRole(id uint) error {
	role := &Response{}
	err := orm.First(database.Instance(), role, scope.IdScope(id))
	if err != nil {
		return err
	}
	if role.Name == GetAdminRoleName() {
		return errors.New("不能操作超级管理员")
	}
	return nil
}

func FindInId(db *gorm.DB, ids []uint) ([]*Response, error) {
	roles := &PageResponse{}
	err := orm.Find(database.Instance(), roles, scope.InIdsScope(ids))
	if err != nil {
		zap_server.ZAPLOG.Error("通过ids查询角色错误", zap.String("错误:", err.Error()))
		return nil, err
	}
	return roles.Item, nil
}

func FindInName(db *gorm.DB, names []string) ([]*Response, error) {
	roles := &PageResponse{}
	err := orm.Find(database.Instance(), roles, scope.InIdsInNamesScopeScope(names))
	if err != nil {
		zap_server.ZAPLOG.Error("通过name查询角色错误", zap.String("错误:", err.Error()))
		return nil, err
	}
	return roles.Item, nil
}

// AddPermForRole
func AddPermForRole(name string, perms [][]string) error {
	oldPerms := casbin.Instance().GetPermissionsForUser(name)
	_, err := casbin.Instance().RemovePolicies(oldPerms)
	if err != nil {
		zap_server.ZAPLOG.Error("add policy err: %+v", zap.String("错误:", err.Error()))
		return err
	}

	if len(perms) == 0 {
		zap_server.ZAPLOG.Debug("没有权限")
		return nil
	}
	var newPerms [][]string
	for _, perm := range perms {
		newPerms = append(newPerms, append([]string{name}, perm...))
	}
	zap_server.ZAPLOG.Info("添加权限到角色", zap_server.Strings("新权限", newPerms))
	_, err = casbin.Instance().AddPolicies(newPerms)
	if err != nil {
		zap_server.ZAPLOG.Error("add policy err: %+v", zap.String("错误:", err.Error()))
		return err
	}

	return nil
}

func GetRoleNames() ([]string, error) {
	var roleNames []string
	err := database.Instance().Model(&Role{}).Select("name").Find(&roleNames).Error
	if err != nil {
		return roleNames, fmt.Errorf("获取角色名称错误 %w", err)
	}
	return roleNames, nil
}
