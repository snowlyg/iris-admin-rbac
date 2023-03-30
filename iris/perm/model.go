package perm

import (
	"errors"

	"github.com/snowlyg/helper/str"
	"github.com/snowlyg/iris-admin/server/zap_server"
	"gorm.io/gorm"
)

const TableName = "permissions"

type PermCollection []Permission

// Permission 权鉴模块
type Permission struct {
	gorm.Model
	BasePermission
}

// BasePermission 权鉴基础模块
type BasePermission struct {
	Name        string `gorm:"index:perm_index,unique;not null ;type:varchar(256)" json:"name" validate:"required,gte=4,lte=50"`
	Act         string `gorm:"index:perm_index;type:varchar(256)" json:"act" validate:"required"`
	DisplayName string `gorm:"type:varchar(256)" json:"displayName"`
	Description string `gorm:"type:varchar(256)" json:"description"`
}

// Create 添加
func (item *Permission) Create(db *gorm.DB) (uint, error) {
	if db == nil {
		return 0, gorm.ErrInvalidDB
	}
	if !CheckNameAndAct(NameScope(item.Name), ActScope(item.Act)) {
		return item.ID, errors.New(str.Join("权限[", item.Name, "-", item.Act, "]已存在"))
	}
	err := db.Model(item).Create(item).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return item.ID, err
	}
	return item.ID, nil
}

// Update 更新
func (item *Permission) Update(db *gorm.DB, scopes ...func(db *gorm.DB) *gorm.DB) error {
	if db == nil {
		return gorm.ErrInvalidDB
	}
	err := db.Model(item).Scopes(scopes...).Updates(item).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return err
	}
	return nil
}

// Delete 删除
func (item *Permission) Delete(db *gorm.DB, scopes ...func(db *gorm.DB) *gorm.DB) error {
	if db == nil {
		return gorm.ErrInvalidDB
	}
	err := db.Model(item).Unscoped().Scopes(scopes...).Delete(item).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return err
	}
	return nil
}
