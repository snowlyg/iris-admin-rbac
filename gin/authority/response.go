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
	err := db.Model(&Authority{}).Scopes(scopes...).First(res).Error
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

	if len(res.Item) > 0 {
		for i := 0; i < len(res.Item); i++ {
			uuid := res.Item[i].Uuid
			res.Item[i].Perms = getPermsForRoleMap(uuid)

			findChildrenAuthority(res.Item[i])
		}
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

	if len(res.Item) > 0 {
		for i := 0; i < len(res.Item); i++ {
			uuid := res.Item[i].Uuid
			res.Item[i].Perms = getPermsForRoleMap(uuid)

			findChildrenAuthority(res.Item[i])
		}

	}

	return nil
}

// findChildrenAuthority
func findChildrenAuthority(item *Response) error {
	err := database.Instance().Where("parent_id = ?", item.Id).Find(&item.Children).Error
	if len(item.Children) > 0 {
		for k := range item.Children {
			err = findChildrenAuthority(&item.Children[k])
		}
	}
	return err
}

// getPermsForRoleMap
func getPermsForRoleMap(uuid string) []map[string]string {
	apisForRoles := []map[string]string{}
	perms := casbin.Instance().GetPermissionsForUser(uuid)
	for _, perm := range perms {
		if len(perm) < 3 {
			continue
		}
		apisForRole := map[string]string{"path": perm[1], "method": perm[2]}
		apisForRoles = append(apisForRoles, apisForRole)
	}
	return apisForRoles
}
