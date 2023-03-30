package api

import (
	"github.com/snowlyg/helper/arr"
	"github.com/snowlyg/iris-admin/server/database/orm"
	"github.com/snowlyg/iris-admin/server/zap_server"
	"gorm.io/gorm"
)

// First 详情
type Response struct {
	orm.Model
	BaseApi
	Children []*Response `json:"children" gorm:"-"`
}

func (res *Response) First(db *gorm.DB, scopes ...func(db *gorm.DB) *gorm.DB) error {
	if db == nil {
		return gorm.ErrInvalidDB
	}
	err := db.Model(&Api{}).Scopes(scopes...).First(res).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return err
	}
	return nil
}

// Paginate 分页
type PageResponse struct {
	Item []*Response
}

func (res *PageResponse) Paginate(db *gorm.DB, pageScope func(db *gorm.DB) *gorm.DB, scopes ...func(db *gorm.DB) *gorm.DB) (int64, error) {
	if db == nil {
		return 0, gorm.ErrInvalidDB
	}
	db = db.Model(&Api{})
	var count int64
	if len(scopes) > 0 {
		db.Scopes(scopes...)
	}
	err := db.Count(&count).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return count, err
	}
	err = db.Scopes(pageScope).Find(&res.Item).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return count, err
	}

	return count, nil
}

func (res *PageResponse) Find(db *gorm.DB, scopes ...func(db *gorm.DB) *gorm.DB) error {
	if db == nil {
		return gorm.ErrInvalidDB
	}
	db = db.Model(&Api{})
	err := db.Scopes(scopes...).Find(&res.Item).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return err
	}

	return nil
}

// FormatApis
func FormatApis(items []*Response) []*Response {
	routers := []*Response{}
	if len(items) > 0 {
		// NOTICE: api,admin,authority,public,oplog
		parentApiCheck := arr.NewCheckArrayType(5)
		for i := 0; i < len(items); i++ {
			if !parentApiCheck.Check(items[i].ApiGroup) {
				router := &Response{}
				router.Path = "/"
				router.Description = items[i].ApiGroup
				routers = append(routers, router)
				parentApiCheck.Add(items[i].ApiGroup)
			}
		}

		for i := 0; i < len(items); i++ {
			for _, router := range routers {
				if router.Description == items[i].ApiGroup {
					router.Children = append(router.Children, items[i])
				}
			}
		}
	}
	return routers
}
