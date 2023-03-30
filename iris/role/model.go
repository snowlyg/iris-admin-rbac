package role

import (
	"github.com/snowlyg/iris-admin/server/zap_server"
	"gorm.io/gorm"
)

type RoleCollection []Request

type Role struct {
	gorm.Model
	BaseRole

	Perms [][]string `gorm:"-" json:"perms"`
}

type BaseRole struct {
	Name        string `gorm:"uniqueIndex;not null; type:varchar(256)" json:"name" validate:"required,gte=4,lte=50" comment:"名称"`
	DisplayName string `gorm:"type:varchar(256)" json:"displayName" comment:"显示名称"`
	Description string `gorm:"type:varchar(256)" json:"description" comment:"描述"`
}

// Create 添加
func (item *Role) Create(db *gorm.DB) (uint, error) {
	if db == nil {
		return 0, gorm.ErrInvalidDB
	}
	err := db.Model(item).Create(item).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return item.ID, err
	}
	return item.ID, nil
}

// Update 更新
func (item *Role) Update(db *gorm.DB, scopes ...func(db *gorm.DB) *gorm.DB) error {
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
func (item *Role) Delete(db *gorm.DB, scopes ...func(db *gorm.DB) *gorm.DB) error {
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
