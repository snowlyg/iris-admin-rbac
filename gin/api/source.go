package api

import (
	"strconv"
	"sync"

	"github.com/gookit/color"
	"github.com/snowlyg/iris-admin/server/database"
	"gorm.io/gorm"
)

var GroupNames = GroupName{
	M: map[string]string{
		"public":    "公共",
		"oplog":     "操作日志",
		"api":       "路由管理",
		"admin":     "管理员管理",
		"authority": "角色管理",
	},
}

type GroupName struct {
	M map[string]string
	sync.Mutex
}

func (gn *GroupName) Add(k, v string) {
	gn.Lock()
	defer gn.Unlock()

	gn.M[k] = v
}

func (gn *GroupName) Get(k string) string {
	gn.Lock()
	defer gn.Unlock()

	return gn.M[k]
}

type source struct {
	routes         []map[string]string
	AuthorityTypes map[string]int
}

// New
func New(routes []map[string]string, authorityTypes map[string]int) *source {
	return &source{
		routes:         routes,
		AuthorityTypes: authorityTypes,
	}
}

func (s *source) GetSources() ApiCollection {
	apis := make(ApiCollection, 0, len(s.routes))
	for _, permRoute := range s.routes {
		group := permRoute["group"]
		if g := GroupNames.Get(group); g != "" {
			group = g
		}
		var isMenu int64 = 0
		isMenu, _ = strconv.ParseInt(permRoute["is_menu"], 10, 64)
		api := Api{BaseApi: BaseApi{
			Path:          permRoute["path"],
			Description:   permRoute["desc"],
			ApiGroup:      group,
			AuthorityType: s.AuthorityTypes[permRoute["path"]],
			Method:        permRoute["method"],
			IsMenu:        isMenu,
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
