package authority

import (
	"github.com/snowlyg/iris-admin/server/casbin"
	"github.com/snowlyg/iris-admin/server/database"
	"github.com/snowlyg/iris-admin/server/database/orm"
	"github.com/snowlyg/iris-admin/server/zap_server"
	"gorm.io/gorm"
)

type Response struct {
	orm.Model
	BaseAuthority
	Uuid     string              `json:"uuid"`
	Children []Response          `json:"children" gorm:"-"`
	Perms    []map[string]string `json:"perms" gorm:"-"`
}

func (res *Response) First(db *gorm.DB, scopes ...func(db *gorm.DB) *gorm.DB) error {
	if db == nil {
		return gorm.ErrInvalidDB
	}
	err := db.Model(&Authority{}).Scopes(scopes...).First(res).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return err
	}

	res.Perms = getPermsForRoleMap(res.Uuid)
	res.findChildrenAuthority()

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
	db = db.Model(&Authority{})
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
	db = db.Model(&Authority{})
	err := db.Scopes(scopes...).Find(&res.Item).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return err
	}

	return nil
}

// findChildrenAuthority
func (item *Response) findChildrenAuthority() error {
	db := database.Instance()
	if db == nil {
		return gorm.ErrInvalidDB
	}
	err := db.Where("parent_id = ?", item.Id).Find(&item.Children).Error
	if len(item.Children) > 0 {
		for k := range item.Children {
			err = item.Children[k].findChildrenAuthority()
		}
	}
	return err
}

// getPermsForRoleMap
func getPermsForRoleMap(uuid string) []map[string]string {
	apisForRoles := []map[string]string{}
	ca := casbin.Instance()
	if ca == nil {
		return nil
	}
	perms := ca.GetPermissionsForUser(uuid)
	for _, perm := range perms {
		if len(perm) < 3 {
			continue
		}
		apisForRole := map[string]string{"path": perm[1], "method": perm[2]}
		apisForRoles = append(apisForRoles, apisForRole)
	}
	return apisForRoles
}
