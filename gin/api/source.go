package api

import (
	"github.com/gookit/color"
	"github.com/snowlyg/iris-admin/server/database"
	"gorm.io/gorm"
)

type source struct {
	routes         []map[string]string
	AuthorityTypes map[string]int
}

func New(routes []map[string]string, authorityTypes map[string]int) *source {
	return &source{
		routes:         routes,
		AuthorityTypes: authorityTypes,
	}
}

func (s *source) GetSources() ApiCollection {
	apis := make(ApiCollection, 0, len(s.routes))
	for _, permRoute := range s.routes {
		api := Api{BaseApi: BaseApi{
			Path:          permRoute["path"],
			Description:   permRoute["desc"],
			ApiGroup:      permRoute["group"],
			AuthorityType: s.AuthorityTypes[permRoute["path"]],
			Method:        permRoute["method"],
		}}
		apis = append(apis, api)
	}
	return apis
}

func (s *source) Init() error {
	if s.GetSources() == nil {
		return nil
	}
	return database.Instance().Transaction(func(tx *gorm.DB) error {
		if err := tx.Unscoped().Where("1 = 1").Delete(&Api{}).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		if err := CreatenInBatches(tx, s.GetSources()); err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> apis 表初始数据成功!")
		return nil
	})
}
