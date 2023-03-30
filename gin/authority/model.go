package authority

import (
	"github.com/snowlyg/iris-admin/server/zap_server"
	"gorm.io/gorm"
)

type AuthorityCollection []Authority

type Authority struct {
	gorm.Model
	BaseAuthority
	Uuid string `json:"uuid" gorm:"uniqueIndex;not null;type:varchar(64);comment:角色标识" binding:"required"`

	Children []Authority `json:"children" gorm:"-"`
	Perms    [][]string  `json:"perms" gorm:"-"`
}

type BaseAuthority struct {
	AuthorityName string `json:"authorityName" gorm:"comment:角色名"`
	AuthorityType int    `json:"authorityType" gorm:"comment:角色类型"`
	ParentId      uint   `json:"parentId" gorm:"default:0;comment:父角色ID"`
	DefaultRouter string `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"`
}

func (item *Authority) mc() map[string]interface{} {
	return map[string]interface{}{
		"authority_name": item.AuthorityName,
		"authority_type": item.AuthorityType,
		"parent_id":      item.ParentId,
		"default_router": item.DefaultRouter,
	}
}

// Create 添加
func (item *Authority) Create(db *gorm.DB) (uint, error) {
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
func (item *Authority) Update(db *gorm.DB, scopes ...func(db *gorm.DB) *gorm.DB) error {
	if db == nil {
		return gorm.ErrInvalidDB
	}
	err := db.Model(item).Scopes(scopes...).Updates(item.mc()).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return err
	}
	return nil
}

// Delete 删除
func (item *Authority) Delete(db *gorm.DB, scopes ...func(db *gorm.DB) *gorm.DB) error {
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
