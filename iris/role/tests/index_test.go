package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/helper/str"
	"github.com/snowlyg/httptest"
	rbac "github.com/snowlyg/iris-admin-rbac/iris"
	"github.com/snowlyg/iris-admin/server/web"
	"github.com/snowlyg/iris-admin/server/web/web_gin/response"
)

var (
	url = "/api/v1/roles" // url
)

func TestList(t *testing.T) {

	TestClient := httptest.Instance(t, str.Join("http://", web.CONFIG.System.Addr), TestServer.GetEngine())
	TestClient.Login(rbac.LoginUrl, nil)
	if TestClient == nil {
		return
	}

	pageKeys := httptest.Responses{
		{Key: "status", Value: http.StatusOK},
		{Key: "message", Value: response.ResponseOkMessage},
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
	TestClient.GET(url, pageKeys, httptest.RequestFunc)
}

func TestCreate(t *testing.T) {
	TestClient := httptest.Instance(t, str.Join("http://", web.CONFIG.System.Addr), TestServer.GetEngine())
	TestClient.Login(rbac.LoginUrl, nil)
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
	TestClient := httptest.Instance(t, str.Join("http://", web.CONFIG.System.Addr), TestServer.GetEngine())
	TestClient.Login(rbac.LoginUrl, nil)
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
		{Key: "status", Value: http.StatusOK},
		{Key: "message", Value: response.ResponseOkMessage},
	}
	TestClient.POST(fmt.Sprintf("%s/%d", url, id), pageKeys, httptest.NewWithJsonParamFunc(update))
}

func TestGetById(t *testing.T) {
	TestClient := httptest.Instance(t, str.Join("http://", web.CONFIG.System.Addr), TestServer.GetEngine())
	TestClient.Login(rbac.LoginUrl, nil)
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
		{Key: "status", Value: http.StatusOK},
		{Key: "message", Value: response.ResponseOkMessage},
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
		{Key: "status", Value: http.StatusOK},
		{Key: "message", Value: response.ResponseOkMessage},
		{Key: "data", Value: httptest.Responses{
			{Key: "id", Value: 1, Type: "ge"},
		},
		},
	}
	return TestClient.POST(url, pageKeys, httptest.NewWithJsonParamFunc(data)).GetId()
}

func Delete(TestClient *httptest.Client, id uint) {
	pageKeys := httptest.Responses{
		{Key: "status", Value: http.StatusOK},
		{Key: "message", Value: response.ResponseOkMessage},
	}
	TestClient.DELETE(fmt.Sprintf("%s/%d", url, id), pageKeys)
}
