package api

import (
	"github.com/snowlyg/iris-admin/server/database/orm"
	"github.com/snowlyg/iris-admin/server/zap_server"
	"gorm.io/gorm"
)

// First 详情
type Response struct {
	orm.Model
	BaseApi
}

func (res *Response) First(db *gorm.DB, scopes ...func(db *gorm.DB) *gorm.DB) error {
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
	db = db.Model(&Api{})
	var count int64
	err := db.Scopes(scopes...).Count(&count).Error
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
	db = db.Model(&Api{})
	err := db.Scopes(scopes...).Find(&res.Item).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return err
	}

	return nil
}
