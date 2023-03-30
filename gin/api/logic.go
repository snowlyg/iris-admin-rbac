package api

import (
	"github.com/snowlyg/iris-admin/g"
	"github.com/snowlyg/iris-admin/server/casbin"
	"github.com/snowlyg/iris-admin/server/database"
	"github.com/snowlyg/iris-admin/server/database/scope"
	"github.com/snowlyg/iris-admin/server/zap_server"
	"gorm.io/gorm"
)

// CreatenInBatches 批量加入
func CreatenInBatches(db *gorm.DB, apis ApiCollection) error {
	if db == nil {
		return gorm.ErrInvalidDB
	}
	err := db.Model(&Api{}).CreateInBatches(&apis, 500).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return err
	}
	return nil
}

func Delete(id uint, req DeleteApiReq) error {
	api := &Api{}
	if err := api.Delete(database.Instance(), scope.IdScope(id)); err != nil {
		return err
	}
	if err := casbin.ClearCasbin(1, req.Path, req.Method); err != nil {
		return err
	}
	return nil
}

func BatcheDelete(ids []uint) error {
	apis := &PageResponse{}
	err := apis.Find(database.Instance(), scope.InIdsScope(ids))
	if err != nil {
		return err
	}

	err = database.Instance().Transaction(func(tx *gorm.DB) error {
		api := &Api{}
		if err := api.Delete(tx, scope.InIdsScope(ids)); err != nil {
			return err
		}
		for _, api := range apis.Item {
			if err := casbin.ClearCasbin(1, api.Path, api.Method); err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// GetApisForRole
func GetApisForRole() (map[int][][]string, error) {
	apis := ApiCollection{}
	err := database.Instance().Model(&Api{}).Where("is_menu=?", g.StatusUnknown).Find(&apis).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return nil, err
	}
	apisForRoles := map[int][][]string{}
	for _, api := range apis {
		apisForRole := []string{api.Path, api.Method}
		apisForRoles[api.AuthorityType] = append(apisForRoles[api.AuthorityType], apisForRole)
	}
	return apisForRoles, nil
}
