package admin

import (
	"path/filepath"
	"regexp"

	"github.com/snowlyg/iris-admin/server/database/orm"
	"github.com/snowlyg/iris-admin/server/web"
	"github.com/snowlyg/iris-admin/server/zap_server"
	"gorm.io/gorm"
)

type Response struct {
	orm.Model
	BaseAdmin
	Authorities []Authorities `gorm:"-" json:"authorities"`
}

type Authorities struct {
	AuthorityName string `json:"authorityName" gorm:"-" `
	Uuid          string `json:"uuid" gorm:"-" `
}

func (res *Response) ToString() {
	if res.Avatar.HeaderImg == "" {
		return
	}
	re := regexp.MustCompile("^http")
	if !re.MatchString(res.Avatar.HeaderImg) {
		res.Avatar.HeaderImg = filepath.ToSlash(web.ToStaticUrl(res.Avatar.HeaderImg))
	}
}

type LoginResponse struct {
	orm.ReqId
	Password     string   `json:"password"`
	AuthorityIds []string `gorm:"-" json:"authorityIds"`
}

func (res *Response) First(db *gorm.DB, scopes ...func(db *gorm.DB) *gorm.DB) error {
	if db == nil {
		return gorm.ErrInvalidDB
	}
	err := db.Model(&Admin{}).Scopes(scopes...).First(res).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return err
	}
	// 查询用户角色
	transform(res)
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
	db = db.Model(&Admin{})
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
	// 查询用户角色
	transform(res.Item...)
	return count, nil
}

func (res *PageResponse) Find(db *gorm.DB, scopes ...func(db *gorm.DB) *gorm.DB) error {
	if db == nil {
		return gorm.ErrInvalidDB
	}
	db = db.Model(&Admin{})
	err := db.Scopes(scopes...).Find(&res.Item).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return err
	}
	// 查询用户角色
	transform(res.Item...)
	return nil
}
