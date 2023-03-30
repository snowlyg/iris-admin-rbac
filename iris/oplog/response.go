package oplog

import (
	"github.com/snowlyg/iris-admin/server/operation"
	"github.com/snowlyg/iris-admin/server/zap_server"
	"gorm.io/gorm"
)

type Response struct {
	operation.Oplog
}

func (res *Response) First(db *gorm.DB, scopes ...func(db *gorm.DB) *gorm.DB) error {
	if db == nil {
		return gorm.ErrInvalidDB
	}
	err := db.Model(&operation.Oplog{}).Scopes(scopes...).First(res).Error
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
	db = db.Model(&operation.Oplog{})
	var count int64
	err := db.Scopes(scopes...).Count(&count).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return 0, err
	}
	err = db.Scopes(pageScope).Find(&res.Item).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return 0, err
	}

	return count, nil
}

func (res *PageResponse) Find(db *gorm.DB, scopes ...func(db *gorm.DB) *gorm.DB) error {
	if db == nil {
		return gorm.ErrInvalidDB
	}
	db = db.Model(&operation.Oplog{})
	err := db.Scopes(scopes...).Find(&res.Item).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return err
	}

	return nil
}
