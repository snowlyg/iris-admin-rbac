package authority

import (
	"errors"

	"github.com/snowlyg/iris-admin/server/casbin"
	"github.com/snowlyg/iris-admin/server/database"
	"github.com/snowlyg/iris-admin/server/database/scope"
	"github.com/snowlyg/iris-admin/server/zap_server"
	"gorm.io/gorm"
)

var ErrRoleNameInvalide = errors.New("角色标识已经被使用")

// GetAdminRoleName 获管理员角色名称
func GetAdminRoleName() string {
	return "admin"
}

// Copy 复制
func Copy(id uint, req *CreateAuthorityRequest) (uint, error) {
	oldAuthority := &Response{}
	err := oldAuthority.First(database.Instance(), scope.IdScope(id))
	if err != nil {
		return 0, err
	}

	if _, err := FindByUuid(AuthorityUuidScope(req.Uuid)); !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, ErrRoleNameInvalide
	}
	authority := &Authority{BaseAuthority: req.BaseAuthority}
	authority.ParentId = oldAuthority.ParentId
	authority.AuthorityType = oldAuthority.AuthorityType
	authority.DefaultRouter = oldAuthority.DefaultRouter
	authority.Uuid = req.Uuid
	newId, err := authority.Create(database.Instance())
	if err != nil {
		return 0, err
	}
	err = AddPermForRole(authority.Uuid, authority.Perms)
	if err != nil {
		return 0, err
	}
	return newId, nil
}

func Update(id uint, req *Authority) error {
	err := database.Instance().Model(&Authority{}).Scopes(scope.IdScope(id)).Updates(req).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return err
	}
	return nil
}

// Create 添加
func Create(req *CreateAuthorityRequest) (uint, error) {
	if _, err := FindByUuid(AuthorityUuidScope(req.Uuid)); !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, ErrRoleNameInvalide
	}
	authority := &Authority{BaseAuthority: req.BaseAuthority, Uuid: req.Uuid}
	id, err := authority.Create(database.Instance())
	if err != nil {
		return 0, err
	}
	err = AddPermForRole(authority.Uuid, authority.Perms)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// FindByUuid
func FindByUuid(scopes ...func(db *gorm.DB) *gorm.DB) (*Response, error) {
	role := &Response{}
	err := role.First(database.Instance(), scopes...)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func IsAdminRole(id uint) error {
	authority := &Response{}
	err := authority.First(database.Instance(), scope.IdScope(id))
	if err != nil {
		return err
	}
	if authority.AuthorityName == GetAdminRoleName() {
		return errors.New("不能操作超级管理员")
	}
	return nil
}

func FindInId(db *gorm.DB, ids []uint) ([]*Response, error) {
	authorities := &PageResponse{}
	err := authorities.Find(database.Instance(), scope.InIdsScope(ids))
	if err != nil {
		return nil, err
	}
	return authorities.Item, nil
}

func FindInUuid(db *gorm.DB, uuids []string) ([]*Response, error) {
	authorities := &PageResponse{}
	err := authorities.Find(database.Instance(), scope.InUuidsScope(uuids))
	if err != nil {
		return nil, err
	}
	return authorities.Item, nil
}

// AddPermForRole
func AddPermForRole(uuid string, perms [][]string) error {
	oldPerms := casbin.Instance().GetPermissionsForUser(uuid)
	_, err := casbin.Instance().RemovePolicies(oldPerms)
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return err
	}

	if len(perms) == 0 {
		zap_server.ZAPLOG.Debug("权限数据为空")
		return nil
	}

	var newPerms [][]string
	for _, perm := range perms {
		newPerms = append(newPerms, append([]string{uuid}, perm...))
	}
	zap_server.ZAPLOG.Info("添加权限到角色:", zap_server.Strings("新权限", newPerms))
	b, err := casbin.Instance().AddPolicies(newPerms)
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return err
	}
	if !b {
		zap_server.ZAPLOG.Error("权限添加失败")
		return errors.New("权限添加失败")
	}

	return nil
}

func GetRoleIds() ([]uint, error) {
	var roleIds []uint
	err := database.Instance().Model(&Authority{}).Select("authority_id").Find(&roleIds).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return roleIds, err
	}
	return roleIds, nil
}
