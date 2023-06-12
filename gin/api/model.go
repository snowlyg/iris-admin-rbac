package api

import (
	"github.com/snowlyg/iris-admin/server/zap_server"
	"gorm.io/gorm"
)

const TableName = "apis"

type ApiCollection []Api

// Api 权鉴模块
type Api struct {
	gorm.Model
	BaseApi
}

// BaseApi 权鉴基础模块
type BaseApi struct {
	Path          string `json:"path" gorm:"comment:api路径" binding:"required"`
	Description   string `json:"description" gorm:"comment:api中文描述" binding:"required"`
	ApiGroup      string `json:"apiGroup" gorm:"comment:api组" binding:"required"`
	Method        string `json:"method" gorm:"default:POST;comment:方法" binding:"required"`
	AuthorityType int    `json:"authorityType" gorm:"comment:角色类型"`
	IsMenu        int64  `gorm:"column:is_menu;type:tinyint(1);not null;default:0" json:"isMenu"` // 是否显示 2:menu, 1为 menu and api，0为 api
}

func (item *Api) mc() map[string]interface{} {
	return map[string]interface{}{
		"is_menu":        item.IsMenu,
		"path":           item.Path,
		"description":    item.Description,
		"api_group":      item.ApiGroup,
		"method":         item.Method,
		"authority_type": item.AuthorityType,
	}
}

// Create 添加
func (item *Api) Create(db *gorm.DB) (uint, error) {
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
func (item *Api) Update(db *gorm.DB, scopes ...func(db *gorm.DB) *gorm.DB) error {
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
func (item *Api) Delete(db *gorm.DB, scopes ...func(db *gorm.DB) *gorm.DB) error {
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
