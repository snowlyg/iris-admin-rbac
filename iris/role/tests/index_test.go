package tests

import (
	"fmt"
	"testing"

	"github.com/snowlyg/httptest"
	rbac "github.com/snowlyg/iris-admin-rbac/iris"
)

var (
	url = "/api/v1/roles" // url
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
					{Key: "name", Value: "SuperAdmin"},
					{Key: "displayName", Value: "超级管理员"},
					{Key: "description", Value: "超级管理员"},
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
		"name":        "test_display_name",
		"displayName": "测试名称",
		"description": "测试描述信息",
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
		"name":        "update_test_display_name",
		"displayName": "测试名称",
		"description": "测试描述信息",
	}
	id := Create(TestClient, data)
	if id == 0 {
		t.Fatalf("测试添加用户失败 id=%d", id)
	}
	defer Delete(TestClient, id)

	update := map[string]interface{}{
		"name":        "update_test_udisplay_name",
		"displayName": "更新测试名称",
		"description": "更新测试描述信息",
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
		"name":        "getbyid_test_display_name",
		"displayName": "更新测试名称",
		"description": "测试描述信息",
	}
	id := Create(TestClient, data)
	if id == 0 {
		t.Fatalf("测试添加失败 id=%d", id)
	}
	defer Delete(TestClient, id)

	pageKeys := httptest.Responses{
		{Key: "code", Value: 2000},
		{Key: "message", Value: "请求成功"},
		{Key: "data", Value: httptest.Responses{
			{Key: "id", Value: 1, Type: "ge"},
			{Key: "name", Value: data["name"].(string)},
			{Key: "displayName", Value: data["displayName"].(string)},
			{Key: "description", Value: data["description"].(string)},
			{Key: "updatedAt", Value: "", Type: "notempty"},
			{Key: "createdAt", Value: "", Type: "notempty"},
			{Key: "createdAt", Value: "", Type: "notempty"},
		},
		},
	}
	TestClient.GET(fmt.Sprintf("%s/%d", url, id), pageKeys)
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
