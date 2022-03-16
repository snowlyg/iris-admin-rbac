package authority

import "gorm.io/gorm"

// AuthorityUuidScope 根据 uuid 查询
// - uuid 名称
func AuthorityUuidScope(uuid string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("uuid = ?", uuid)
	}
}

// AuthorityTypeScope 根据 type 查询
// - authorityType 角色类型
func AuthorityTypeScope(authorityType int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("authority_type = ?", authorityType)
	}
}
