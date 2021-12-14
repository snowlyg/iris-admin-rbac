package authority

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func GetMigration() *gormigrate.Migration {
	return &gormigrate.Migration{
		ID: "20211214120700_create_authorities_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&Authority{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("authorities")
		},
	}
}
