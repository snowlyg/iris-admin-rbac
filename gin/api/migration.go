package api

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/snowlyg/helper/global"
	"github.com/snowlyg/helper/str"
	"gorm.io/gorm"
)

func GetMigration() *gormigrate.Migration {
		id :=	time.Now().Format(global.TimestampLayout)
		return  &gormigrate.Migration{
			ID: str.Join(id,"_","create_apis_table"),
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&Api{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("apis")
			},
		}
}
