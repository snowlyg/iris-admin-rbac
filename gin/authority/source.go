package authority

import (
	"github.com/gookit/color"
	"github.com/snowlyg/iris-admin-rbac/gin/api"
	"github.com/snowlyg/iris-admin/server/database"
	"github.com/snowlyg/iris-admin/server/database/orm"
	"github.com/snowlyg/multi"
	"gorm.io/gorm"
)

var Source = new(source)

type source struct{}

func GetSources() ([]*Authority, error) {
	apis, err := api.GetApisForRole()
	if err != nil {
		return nil, err
	}

	sources := []*Authority{
		{
			BaseAuthority: BaseAuthority{
				AuthorityName: "超级管理员",
				AuthorityType: multi.AdminAuthority,
				ParentId:      0,
				DefaultRouter: "",
			},
			Uuid:  "super_admin",
			Perms: apis[multi.AdminAuthority],
		},
		{
			Uuid: "tenancy_admin",
			BaseAuthority: BaseAuthority{
				AuthorityName: "商户管理员",
				AuthorityType: multi.TenancyAuthority,
				ParentId:      0,
				DefaultRouter: "",
			},
			Perms: apis[multi.TenancyAuthority],
		},
		{
			Uuid: "mini_admin",
			BaseAuthority: BaseAuthority{
				AuthorityName: "小程序用户",
				AuthorityType: multi.GeneralAuthority,
				ParentId:      0,
				DefaultRouter: "",
			},
			Perms: apis[multi.GeneralAuthority],
		},
		{
			Uuid: "device_admin",
			BaseAuthority: BaseAuthority{
				AuthorityName: "设备用户",
				AuthorityType: multi.GeneralAuthority,
				ParentId:      0,
				DefaultRouter: "",
			},
			Perms: apis[multi.GeneralAuthority],
		},
	}

	return sources, nil
}

func (s *source) Init() error {
	return database.Instance().Transaction(func(tx *gorm.DB) error {
		if tx.Model(&Authority{}).Where("authority_id IN ?", []int{1}).Find(&[]Authority{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> authotities 表的初始数据已存在!")
			return nil
		}
		sources, err := GetSources()
		if err != nil {
			return err
		}
		for _, source := range sources {
			_, err := orm.Create(database.Instance(), source)
			if err != nil {
				return err
			}
			err = AddPermForRole(source.Uuid, source.Perms)
			if err != nil {
				return err
			}
		}

		color.Info.Println("\n[Mysql] --> authotities 表初始数据成功!")
		return nil
	})
}
