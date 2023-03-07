package authority

import (
	"github.com/snowlyg/iris-admin-rbac/gin/api"
	"github.com/snowlyg/iris-admin/server/database/orm"
	"github.com/snowlyg/iris-admin/server/zap_server"
	"gorm.io/gorm"
)

type Response struct {
	orm.Model
	BaseAuthority
	Uuid string `json:"uuid"`
	Menus    []BaseMenu  `json:"menus" gorm:"many2many:authority_menus;"`
	Children []Authority `json:"children" gorm:"-"`
	Perms    []map[string]string  `json:"perms" gorm:"-"`
}

func (res *Response) First(db *gorm.DB, scopes ...func(db *gorm.DB) *gorm.DB) error {
	err := db.Model(&Authority{}).Scopes(scopes...).First(res).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return err
	}
	apis, err := api.GetApisForRoleMap()
	if err != nil {
		return  err
	}
	res.Perms = apis[res.AuthorityType]
	return nil
}

// Paginate 分页
type PageResponse struct {
	Item []*Response
}

func (res *PageResponse) Paginate(db *gorm.DB, pageScope func(db *gorm.DB) *gorm.DB, scopes ...func(db *gorm.DB) *gorm.DB) (int64, error) {
	db = db.Model(&Authority{})
	var count int64
		if len(scopes)>0{
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

	apis, err := api.GetApisForRoleMap()
	if err != nil {
		return  count,err
	}
	for i := 0; i < len(res.Item); i++ {
		res.Item[i].Perms = apis[res.Item[i].AuthorityType]
	}

	return count, nil
}

func (res *PageResponse) Find(db *gorm.DB, scopes ...func(db *gorm.DB) *gorm.DB) error {
	db = db.Model(&Authority{})
	err := db.Scopes(scopes...).Find(&res.Item).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return err
	}

	
	apis, err := api.GetApisForRoleMap()
	if err != nil {
		return err
	}
	for i := 0; i < len(res.Item); i++ {
		res.Item[i].Perms = apis[res.Item[i].AuthorityType]
	}

	return nil
}
