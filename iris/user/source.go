package user

import (
	"github.com/gookit/color"
	"github.com/snowlyg/iris-admin-rbac/iris/role"
	"github.com/snowlyg/iris-admin/server/database"
	"gorm.io/gorm"
)

var Source = new(source)

type source struct{}

func GetSources() ([]*Request, error) {
	roleNames, err := role.GetRoleNames()
	if err != nil {
		return []*Request{}, err
	}
	var users []*Request
	users = append(users, &Request{
		BaseUser: BaseUser{
			Name:     "超级管理员",
			Username: "admin",
			Intro:    "超级管理员",
			Avatar:   "/images/avatar.jpg",
		},
		Password:  "123456",
		RoleNames: roleNames,
	})
	return users, nil
}

func (s *source) Init() error {
	return database.Instance().Transaction(func(tx *gorm.DB) error {
		if tx.Model(&User{}).Where("id IN ?", []int{1}).Find(&[]User{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> users 表的初始数据已存在!")
			return nil
		}
		sources, err := GetSources()
		if err != nil {
			return err
		}
		for _, source := range sources {
			if _, err := Create(source); err != nil { // 遇到错误时回滚事务
				return err
			}
		}
		color.Info.Println("\n[Mysql] --> users 表初始数据成功!")
		return nil
	})
}
