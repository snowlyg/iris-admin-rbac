package authority

import (
	"github.com/snowlyg/helper/str"
	"gorm.io/gorm"
)

// AuthorityUuidScope
func AuthorityUuidScope(uuid string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("uuid = ?", uuid)
	}
}

// AuthorityTypeScope
func AuthorityTypeScope(authorityType int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("authority_type = ?", authorityType)
	}
}

// AuthorityNameScope
func AuthorityNameScope(authorityName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("authority_name LIKE ?", str.Join(authorityName, "%"))
	}
}

// ParentIdScope
func ParentIdScope(id uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("parent_id = ?", id)
	}
}
