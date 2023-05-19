package admin

import (
	"github.com/snowlyg/helper/str"
	"gorm.io/gorm"
)

// UserNameScope
func UserNameScope(username string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("username = ?", username)
	}
}

// SearchKeyScope
func SearchKeyScope(key string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		k := str.Join(key, "%")
		return db.Where("username Like ? OR nick_name Like ? OR phone Like ?", k, k, k)
	}
}

// StatusScope
func StatusScope(status string) func(db *gorm.DB) *gorm.DB {
	if status != "1" {
		status = "0"
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", status)
	}
}
