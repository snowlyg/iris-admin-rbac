package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/helper/tests"
	rbac "github.com/snowlyg/iris-admin-rbac/gin"
	"github.com/snowlyg/iris-admin/g"
	"github.com/snowlyg/iris-admin/server/web"
	"github.com/snowlyg/iris-admin/server/web/web_gin/response"
)

var (
	url = "/api/v1/admin"
)

func TestList(t *testing.T) {
	if TestServer == nil {
		t.Error("测试服务初始化失败")
		return
	}

	TestClient = TestServer.GetTestLogin(t, rbac.LoginUrl, rbac.LoginResponse)
	if TestClient == nil {
		return
	}
	pageKeys := tests.Responses{
		{Key: "status", Value: http.StatusOK},
		{Key: "message", Value: response.ResponseOkMessage},
		{Key: "data", Value: tests.Responses{
			{Key: "pageSize", Value: 10},
			{Key: "page", Value: 1},
			{Key: "list", Value: []tests.Responses{
				{
					{Key: "id", Value: 1, Type: "ge"},
					{Key: "nickName", Value: "超级管理员"},
					{Key: "username", Value: "admin"},
					{Key: "headerImg", Value: "http://qmplusimg.henrongyi.top/head.png"},
					{Key: "status", Value: g.StatusTrue},
					{Key: "isShow", Value: g.StatusFalse},
					{Key: "phone", Value: "13800138000"},
					{Key: "email", Value: "admin@admin.com"},
					{Key: "authorities", Value: []string{"超级管理员"}},
					{Key: "updatedAt", Value: "", Type: "notempty"},
					{Key: "createdAt", Value: "", Type: "notempty"},
				},
			}},
			{Key: "total", Value: 0, Type: "ge"},
		}},
	}
	TestClient.GET(fmt.Sprintf("%s/getAll", url), pageKeys, tests.RequestParams)
}

func TestCreate(t *testing.T) {
	if TestServer == nil {
		t.Error("测试服务初始化失败")
		return
	}

	client := TestServer.GetTestLogin(t, rbac.LoginUrl, rbac.LoginResponse)
	if client == nil {
		return
	}

	data := map[string]interface{}{
		"nickName":     "测试名称",
		"username":     "create_test_username",
		"authorityIds": []uint{web.AdminAuthorityId},
		"email":        "get@admin.com",
		"phone":        "13800138001",
		"password":     "123456",
	}
	id := Create(client, data)
	if id == 0 {
		t.Fatalf("测试添加用户失败 id=%d", id)
	}
	defer Delete(client, id)
}

func TestUpdate(t *testing.T) {
	if TestServer == nil {
		t.Error("测试服务初始化失败")
		return
	}

	client := TestServer.GetTestLogin(t, rbac.LoginUrl, rbac.LoginResponse)
	if client == nil {
		return
	}
	data := map[string]interface{}{
		"nickName":     "测试名称",
		"username":     "create_test_username_for_update",
		"authorityIds": []uint{web.AdminAuthorityId},
		"email":        "get@admin.com",
		"phone":        "13800138001",
		"password":     "123456",
	}
	id := Create(client, data)
	if id == 0 {
		t.Fatalf("测试添加用户失败 id=%d", id)
	}
	defer Delete(client, id)

	update := map[string]interface{}{
		"nickName": "测试名称",
		"email":    "get@admin.com",
		"phone":    "13800138003",
		"password": "123456",
	}

	pageKeys := tests.Responses{
		{Key: "status", Value: http.StatusOK},
		{Key: "message", Value: response.ResponseOkMessage},
	}
	client.PUT(fmt.Sprintf("%s/updateAdmin/%d", url, id), pageKeys, update)
}

func TestGetById(t *testing.T) {
	if TestServer == nil {
		t.Error("测试服务初始化失败")
		return
	}

	client := TestServer.GetTestLogin(t, rbac.LoginUrl, rbac.LoginResponse)
	if client == nil {
		return
	}
	data := map[string]interface{}{
		"nickName":     "测试名称",
		"username":     "create_test_username_for_get",
		"email":        "get@admin.com",
		"phone":        "13800138001",
		"authorityIds": []uint{web.AdminAuthorityId},
		"password":     "123456",
	}
	id := Create(client, data)
	if id == 0 {
		t.Fatalf("测试添加用户失败 id=%d", id)
	}
	defer Delete(client, id)
	pageKeys := tests.Responses{
		{Key: "status", Value: http.StatusOK},
		{Key: "message", Value: response.ResponseOkMessage},
		{Key: "data", Value: tests.Responses{
			{Key: "id", Value: 1, Type: "ge"},
			{Key: "nickName", Value: data["nickName"].(string)},
			{Key: "username", Value: data["username"].(string)},
			{Key: "status", Value: g.StatusTrue},
			{Key: "email", Value: data["email"].(string)},
			{Key: "phone", Value: data["phone"].(string)},
			{Key: "isShow", Value: g.StatusTrue},
			{Key: "headerImg", Value: "http://qmplusimg.henrongyi.top/head.png"},
			{Key: "updatedAt", Value: "", Type: "notempty"},
			{Key: "createdAt", Value: "", Type: "notempty"},
			{Key: "createdAt", Value: "", Type: "notempty"},
			{Key: "authorities", Value: []string{"超级管理员"}},
		},
		},
	}
	client.GET(fmt.Sprintf("%s/getAdmin/%d", url, id), pageKeys)
}

func Create(client *tests.Client, data map[string]interface{}) uint {
	pageKeys := tests.Responses{
		{Key: "status", Value: http.StatusOK},
		{Key: "message", Value: response.ResponseOkMessage},
		{Key: "data", Value: tests.Responses{
			{Key: "id", Value: 1, Type: "ge"},
		},
		},
	}
	return client.POST(fmt.Sprintf("%s/createAdmin", url), pageKeys, data).GetId()
}

func Delete(client *tests.Client, id uint) {
	pageKeys := tests.Responses{
		{Key: "status", Value: http.StatusOK},
		{Key: "message", Value: response.ResponseOkMessage},
	}
	client.DELETE(fmt.Sprintf("%s/deleteAdmin/%d", url, id), pageKeys)
}
