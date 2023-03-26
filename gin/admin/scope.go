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
		return db.Where(db.Where("username Like ?", str.Join(key, "%")).Or("nick_name Like ?", str.Join(key, "%")).Or("phone Like ?", str.Join(key, "%")))
	}
}

//StatusScope
func StatusScope(status string) func(db *gorm.DB) *gorm.DB {
	if status != "1" {
		status = "0"
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", status)
	}
}
