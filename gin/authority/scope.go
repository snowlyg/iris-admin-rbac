package authority

import "gorm.io/gorm"

// AuthorityNameScope 根据 name 查询
// - name 名称
func AuthorityNameScope(name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("authority_name = ?", name)
	}
}

// AuthorityTypeScope 根据 type 查询
// - authorityType 角色类型
func AuthorityTypeScope(authorityType int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("authority_type = ?", authorityType)
	}
}
