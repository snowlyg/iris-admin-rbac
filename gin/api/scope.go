package api

import (
	"github.com/snowlyg/helper/str"
	"gorm.io/gorm"
)

// AuthorityTypeScope
func AuthorityTypeScope(authorityType int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("authority_type = ?", authorityType)
	}
}

// MethodScope
func MethodScope(method string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("method = ?", method)
	}
}

// ApiGroupScope
func ApiGroupScope(apiGroup string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("api_group = ?", apiGroup)
	}
}

// PathScope
func PathScope(path string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("path LIKE ?", str.Join(path, `%`))
	}
}
