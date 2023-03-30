package perm

import (
	"errors"

	"github.com/snowlyg/iris-admin/server/database"
	"github.com/snowlyg/iris-admin/server/zap_server"
	"gorm.io/gorm"
)

// CreatenInBatches 批量加入
func CreatenInBatches(db *gorm.DB, perms PermCollection) error {
	if db == nil {
		return gorm.ErrInvalidDB
	}
	err := db.Model(&Permission{}).CreateInBatches(&perms, 500).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return err
	}
	return nil
}

// CheckNameAndAct 检测权限是否存在
func CheckNameAndAct(scopes ...func(db *gorm.DB) *gorm.DB) bool {
	perm := &Response{}
	err := perm.First(database.Instance(), scopes...)
	isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	if !isNotFound {
		zap_server.ZAPLOG.Error(err.Error())
	}
	return isNotFound
}

// GetPermsForRole
func GetPermsForRole() ([][]string, error) {
	var permsForRoles [][]string
	perms := PermCollection{}
	err := database.Instance().Model(&Permission{}).Find(&perms).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return nil, err
	}
	for _, perm := range perms {
		permsForRole := []string{perm.Name, perm.Act}
		permsForRoles = append(permsForRoles, permsForRole)
	}
	return permsForRoles, nil
}
