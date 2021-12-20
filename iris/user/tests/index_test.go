package tests

import (
	"fmt"
	"testing"

	"github.com/snowlyg/httptest"
	rbac "github.com/snowlyg/iris-admin-rbac/iris"
	"github.com/snowlyg/iris-admin/server/web/web_iris"
)

var (
	url = "/api/v1/users"
)

func TestList(t *testing.T) {
	if TestServer == nil {
		t.Errorf("测试服务初始化失败")
	}
	TestClient = TestServer.GetTestLogin(t, rbac.LoginUrl, nil)
	if TestClient == nil {
		return
	}

	pageKeys := httptest.Responses{
		{Key: "code", Value: 2000},
		{Key: "message", Value: "请求成功"},
		{Key: "data", Value: httptest.Responses{
			{Key: "pageSize", Value: 10},
			{Key: "page", Value: 1},
			{Key: "items", Value: []httptest.Responses{
				{
					{Key: "id", Value: 1, Type: "ge"},
					{Key: "name", Value: "超级管理员"},
					{Key: "username", Value: "admin"},
					{Key: "intro", Value: "超级管理员"},
					{Key: "avatar", Value: web_iris.ToStaticUrl("/images/avatar.jpg")},
					{Key: "roles", Value: []string{"超级管理员"}},
					{Key: "updatedAt", Value: "", Type: "notempty"},
					{Key: "createdAt", Value: "", Type: "notempty"},
				},
			}},
			{Key: "total", Value: 0, Type: "ge"},
		}},
	}
	TestClient.GET(url, pageKeys, httptest.RequestParams)
}

func TestCreate(t *testing.T) {
	TestClient = TestServer.GetTestLogin(t, rbac.LoginUrl, nil)
	if TestClient == nil {
		return
	}

	data := map[string]interface{}{
		"name":     "测试名称",
		"username": "create_test_username",
		"intro":    "测试描述信息",
		"avatar":   "",
		"password": "123456",
	}
	id := Create(TestClient, data)
	if id == 0 {
		t.Fatalf("测试添加用户失败 id=%d", id)
	}
	defer Delete(TestClient, id)
}

func TestUpdate(t *testing.T) {
	TestClient = TestServer.GetTestLogin(t, rbac.LoginUrl, nil)
	if TestClient == nil {
		return
	}

	data := map[string]interface{}{
		"name":     "测试名称",
		"username": "update_test_username",
		"intro":    "测试描述信息",
		"avatar":   "",
		"password": "123456",
	}
	id := Create(TestClient, data)
	if id == 0 {
		t.Fatalf("测试添加用户失败 id=%d", id)
	}
	defer Delete(TestClient, id)

	update := map[string]interface{}{
		"name":     "更新测试名称",
		"username": "update_test_username",
		"intro":    "更新测试描述信息",
		"avatar":   "",
		"password": "123456",
	}

	pageKeys := httptest.Responses{
		{Key: "code", Value: 2000},
		{Key: "message", Value: "请求成功"},
	}
	TestClient.POST(fmt.Sprintf("%s/%d", url, id), pageKeys, update)
}

func TestGetById(t *testing.T) {
	TestClient = TestServer.GetTestLogin(t, rbac.LoginUrl, nil)
	if TestClient == nil {
		return
	}

	data := map[string]interface{}{
		"name":     "测试名称",
		"username": "getbyid_test_username",
		"intro":    "测试描述信息",
		"avatar":   "",
		"password": "123456",
	}
	id := Create(TestClient, data)
	if id == 0 {
		t.Fatalf("测试添加用户失败 id=%d", id)
	}
	defer Delete(TestClient, id)

	pageKeys := httptest.Responses{
		{Key: "code", Value: 2000},
		{Key: "message", Value: "请求成功"},
		{Key: "data", Value: httptest.Responses{
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
	TestClient.GET(fmt.Sprintf("%s/%d", url, id), pageKeys)
}
func TestChangeAvatar(t *testing.T) {
	TestClient = TestServer.GetTestLogin(t, rbac.LoginUrl, nil)
	if TestClient == nil {
		return
	}
	data := map[string]interface{}{
		"avatar": "/avatar.png",
	}
	pageKeys := httptest.Responses{
		{Key: "code", Value: 2000},
		{Key: "message", Value: "请求成功"},
	}
	TestClient.POST(fmt.Sprintf("%s/changeAvatar", url), pageKeys, data)

	profile := httptest.Responses{
		{Key: "code", Value: 2000},
		{Key: "message", Value: "请求成功"},
		{Key: "data", Value: httptest.Responses{
			{Key: "id", Value: 1, Type: "ge"},
			{Key: "name", Value: "超级管理员"},
			{Key: "username", Value: "admin"},
			{Key: "intro", Value: "超级管理员"},
			{Key: "avatar", Value: web_iris.ToStaticUrl("/avatar.png")},
			{Key: "roles", Value: []string{"超级管理员"}},
			{Key: "updatedAt", Value: "", Type: "notempty"},
			{Key: "createdAt", Value: "", Type: "notempty"},
		},
		},
	}
	TestClient.GET(fmt.Sprintf("%s/profile", url), profile)
}

func Create(TestClient *httptest.Client, data map[string]interface{}) uint {
	pageKeys := httptest.Responses{
		{Key: "code", Value: 2000},
		{Key: "message", Value: "请求成功"},
		{Key: "data", Value: httptest.Responses{
			{Key: "id", Value: 1, Type: "ge"},
		},
		},
	}
	return TestClient.POST(url, pageKeys, data).GetId()
}

func Delete(TestClient *httptest.Client, id uint) {
	pageKeys := httptest.Responses{
		{Key: "code", Value: 2000},
		{Key: "message", Value: "请求成功"},
	}
	TestClient.DELETE(fmt.Sprintf("%s/%d", url, id), pageKeys)
}
