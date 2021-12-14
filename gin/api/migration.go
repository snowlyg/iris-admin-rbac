package api

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func GetMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20211214120700_create_apis_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&Api{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("apis")
		},
	}
}
