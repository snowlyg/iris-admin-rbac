package user

import (
	"github.com/snowlyg/iris-admin/server/zap_server"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	BaseUser
	Password  string   `gorm:"type:varchar(250)" json:"password" validate:"required"`
	RoleNames []string `gorm:"-" json:"role_names"`
}

type BaseUser struct {
	Name     string `gorm:"index;not null; type:varchar(60)" json:"name"`
	Username string `gorm:"uniqueIndex;not null;type:varchar(60)" json:"username" validate:"required"`
	Intro    string `gorm:"not null; type:varchar(512)" json:"intro"`
	Avatar   string `gorm:"type:varchar(1024)" json:"avatar"`
}

type Avatar struct {
	Avatar string `json:"avatar"`
}

// Create 添加
func (item *User) Create(db *gorm.DB) (uint, error) {
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
func (item *User) Update(db *gorm.DB, scopes ...func(db *gorm.DB) *gorm.DB) error {
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
func (item *User) Delete(db *gorm.DB, scopes ...func(db *gorm.DB) *gorm.DB) error {
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
