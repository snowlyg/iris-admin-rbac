package api

import (
	"fmt"

	"github.com/snowlyg/iris-admin/server/casbin"
	"github.com/snowlyg/iris-admin/server/database"
	"github.com/snowlyg/iris-admin/server/database/orm"
	"github.com/snowlyg/iris-admin/server/database/scope"
	"github.com/snowlyg/iris-admin/server/zap_server"
	"gorm.io/gorm"
)

// CreatenInBatches 批量加入
func CreatenInBatches(db *gorm.DB, apis ApiCollection) error {
	err := db.Model(&Api{}).CreateInBatches(&apis, 500).Error
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return err
	}
	return nil
}

func Delete(id uint, req DeleteApiReq) error {
	if err := orm.Delete(database.Instance(), &Api{}, scope.IdScope(id)); err != nil {
		return err
	}
	if err := casbin.ClearCasbin(1, req.Path, req.Method); err != nil {
		return err
	}
	return nil
}

func BatcheDelete(ids []uint) error {
	apis := &PageResponse{}
	err := orm.Find(database.Instance(), apis, scope.InIdsScope(ids))
	if err != nil {
		return err
	}

	err = database.Instance().Transaction(func(tx *gorm.DB) error {
		if err := orm.Delete(tx, &Api{}, scope.InIdsScope(ids)); err != nil {
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
	err := database.Instance().Model(&Api{}).Find(&apis).Error
	if err != nil {
		return nil, fmt.Errorf("获取权限错误 %w", err)
	}
	apisForRoles := map[int][][]string{}
	for _, api := range apis {
		apisForRole := []string{api.Path, api.Method}
		apisForRoles[api.AuthorityType] = append(apisForRoles[api.AuthorityType], apisForRole)
	}
	return apisForRoles, nil
}
