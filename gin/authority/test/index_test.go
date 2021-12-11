package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/helper/tests"
	rbac "github.com/snowlyg/iris-admin-rbac/gin"
	"github.com/snowlyg/iris-admin/server/web/web_gin/response"
	"github.com/snowlyg/multi"
)

var (
	url = "/api/v1/authority" // url
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
			{Key: "items", Value: []tests.Responses{
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
	TestClient.GET(url, pageKeys, tests.RequestParams)
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
		"authorityId":   "999",
		"authorityName": "测试角色",
		"parentId":      "0",
		"authorityType": multi.AdminAuthority,
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
		"authorityId":   "999",
		"authorityName": "测试角色",
		"parentId":      "0",
		"authorityType": multi.AdminAuthority,
	}
	id := Create(client, data)
	if id == 0 {
		t.Fatalf("测试添加用户失败 id=%d", id)
	}
	defer Delete(client, id)

	update := map[string]interface{}{
		"authorityId":   "999",
		"authorityName": "测试角色",
		"parentId":      "0",
		"authorityType": multi.AdminAuthority,
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

	client := TestServer.GetTestLogin(t, rbac.LoginUrl, rbac.LoginResponse)
	if client == nil {
		return
	}
	data := map[string]interface{}{
		"authorityId":   "999",
		"authorityName": "测试角色",
		"parentId":      "0",
		"authorityType": multi.AdminAuthority,
	}
	id := Create(client, data)
	if id == 0 {
		t.Fatalf("测试添加失败 id=%d", id)
	}
	defer Delete(client, id)

	pageKeys := tests.Responses{
		{Key: "status", Value: http.StatusOK},
		{Key: "message", Value: response.ResponseOkMessage},
		{Key: "data", Value: tests.Responses{
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
	return client.POST(url, pageKeys, data).GetId()
}

func Delete(client *tests.Client, id uint) {
	pageKeys := tests.Responses{
		{Key: "status", Value: http.StatusOK},
		{Key: "message", Value: response.ResponseOkMessage},
	}
	client.DELETE(fmt.Sprintf("%s/%d", url, id), pageKeys)
}
