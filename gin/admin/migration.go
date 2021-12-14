package admin

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func GetMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20211214120700_create_admins_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&Admin{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("admins")
		},
	}
}
