package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/helper/tests"
	v1 "github.com/snowlyg/iris-admin-rbac/gin"
	"github.com/snowlyg/iris-admin/g"
	"github.com/snowlyg/iris-admin/server/web/web_gin/response"
)

var (
	loginUrl = "/api/v1/public/admin/login"

	url = "/api/v1/admin"
)

func TestList(t *testing.T) {
	if TestServer == nil {
		t.Error("测试服务初始化失败")
		return
	}

	TestClient = TestServer.GetTestLogin(t, loginUrl, v1.LoginResponse)
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
					{Key: "is_show", Value: g.StatusFalse},
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

	client := TestServer.GetTestLogin(t, loginUrl, v1.LoginResponse)
	if client == nil {
		return
	}

	data := map[string]interface{}{
		"nickName":     "测试名称",
		"username":     "create_test_username",
		"intro":        "测试描述信息",
		"authorityIds": []uint{1},
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

	client := TestServer.GetTestLogin(t, loginUrl, v1.LoginResponse)
	if client == nil {
		return
	}
	data := map[string]interface{}{
		"name":     "测试名称",
		"username": "update_test_username",
		"intro":    "测试描述信息",
		"avatar":   "",
		"password": "123456",
	}
	id := Create(client, data)
	if id == 0 {
		t.Fatalf("测试添加用户失败 id=%d", id)
	}
	defer Delete(client, id)

	update := map[string]interface{}{
		"name":     "更新测试名称",
		"username": "update_test_username",
		"intro":    "更新测试描述信息",
		"avatar":   "",
		"password": "123456",
	}

	pageKeys := tests.Responses{
		{Key: "status", Value: http.StatusOK},
		{Key: "message", Value: response.ResponseOkMessage},
	}
	client.POST(fmt.Sprintf("%s/%d", url, id), pageKeys, update)
}

func TestGetById(t *testing.T) {
	if TestServer == nil {
		t.Error("测试服务初始化失败")
		return
	}

	client := TestServer.GetTestLogin(t, loginUrl, v1.LoginResponse)
	if client == nil {
		return
	}
	data := map[string]interface{}{
		"name":     "测试名称",
		"username": "getbyid_test_username",
		"intro":    "测试描述信息",
		"avatar":   "",
		"password": "123456",
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
			{Key: "name", Value: data["name"].(string)},
			{Key: "username", Value: data["username"].(string)},
			{Key: "intro", Value: data["intro"].(string)},
			{Key: "avatar", Value: data["avatar"].(string)},
			{Key: "updatedAt", Value: "", Type: "notempty"},
			{Key: "createdAt", Value: "", Type: "notempty"},
			{Key: "createdAt", Value: "", Type: "notempty"},
			{Key: "roles", Value: []string{}, Type: "null"},
		},
		},
	}
	client.GET(fmt.Sprintf("%s/%d", url, id), pageKeys)
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
	return client.POST(fmt.Sprintf("%s/%s", url, "create"), pageKeys, data).GetId()
}

func Delete(client *tests.Client, id uint) {
	pageKeys := tests.Responses{
		{Key: "status", Value: http.StatusOK},
		{Key: "message", Value: response.ResponseOkMessage},
	}
	client.DELETE(fmt.Sprintf("%s/%d", url, id), pageKeys)
}
