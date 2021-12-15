package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/httptest"
	rbac "github.com/snowlyg/iris-admin-rbac/gin"
	"github.com/snowlyg/iris-admin/server/web"
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
	pageKeys := httptest.Responses{
		{Key: "status", Value: http.StatusOK},
		{Key: "message", Value: response.ResponseOkMessage},
		{Key: "data", Value: httptest.Responses{
			{Key: "pageSize", Value: 10},
			{Key: "page", Value: 1},
			{Key: "list", Value: []httptest.Responses{
				{
					{Key: "id", Value: web.DeviceAuthorityId, Type: "ge"},
					{Key: "authorityName", Value: "设备用户"},
					{Key: "authorityType", Value: multi.GeneralAuthority},
					{Key: "parentId", Value: 0},
					{Key: "defaultRouter", Value: "dashboard"},
					{Key: "updatedAt", Value: "", Type: "notempty"},
					{Key: "createdAt", Value: "", Type: "notempty"},
				},
				{
					{Key: "id", Value: web.LiteAuthorityId, Type: "ge"},
					{Key: "authorityName", Value: "小程序用户"},
					{Key: "authorityType", Value: multi.GeneralAuthority},
					{Key: "parentId", Value: 0},
					{Key: "defaultRouter", Value: "dashboard"},
					{Key: "updatedAt", Value: "", Type: "notempty"},
					{Key: "createdAt", Value: "", Type: "notempty"},
				},
				{
					{Key: "id", Value: web.TenancyAuthorityId, Type: "ge"},
					{Key: "authorityName", Value: "商户管理员"},
					{Key: "authorityType", Value: multi.TenancyAuthority},
					{Key: "parentId", Value: 0},
					{Key: "defaultRouter", Value: "dashboard"},
					{Key: "updatedAt", Value: "", Type: "notempty"},
					{Key: "createdAt", Value: "", Type: "notempty"},
				},
				{
					{Key: "id", Value: web.AdminAuthorityId, Type: "ge"},
					{Key: "authorityName", Value: "超级管理员"},
					{Key: "authorityType", Value: multi.AdminAuthority},
					{Key: "parentId", Value: 0},
					{Key: "defaultRouter", Value: "dashboard"},
					{Key: "updatedAt", Value: "", Type: "notempty"},
					{Key: "createdAt", Value: "", Type: "notempty"},
				},
			}},
			{Key: "total", Value: 0, Type: "ge"},
		}},
	}
	requestParams := map[string]interface{}{"page": 1, "pageSize": 10, "orderBy": "id"}
	TestClient.GET(fmt.Sprintf("%s/getAuthorityList", url), pageKeys, requestParams)
}

func TestGetAdminAuthorityList(t *testing.T) {
	if TestServer == nil {
		t.Error("测试服务初始化失败")
		return
	}

	TestClient = TestServer.GetTestLogin(t, rbac.LoginUrl, rbac.LoginResponse)
	if TestClient == nil {
		return
	}
	pageKeys := httptest.Responses{
		{Key: "status", Value: http.StatusOK},
		{Key: "message", Value: response.ResponseOkMessage},
		{Key: "data", Value: httptest.Responses{
			{Key: "pageSize", Value: 10},
			{Key: "page", Value: 1},
			{Key: "list", Value: []httptest.Responses{
				{
					{Key: "id", Value: web.AdminAuthorityId, Type: "ge"},
					{Key: "authorityName", Value: "超级管理员"},
					{Key: "authorityType", Value: multi.AdminAuthority},
					{Key: "parentId", Value: 0},
					{Key: "defaultRouter", Value: "dashboard"},
					{Key: "updatedAt", Value: "", Type: "notempty"},
					{Key: "createdAt", Value: "", Type: "notempty"},
				},
			}},
			{Key: "total", Value: 0, Type: "ge"},
		}},
	}
	requestParams := map[string]interface{}{"page": 1, "pageSize": 10, "orderBy": "id"}
	TestClient.GET(fmt.Sprintf("%s/getAdminAuthorityList", url), pageKeys, requestParams)
}

func TestGetTenancyAuthorityList(t *testing.T) {
	if TestServer == nil {
		t.Error("测试服务初始化失败")
		return
	}

	TestClient = TestServer.GetTestLogin(t, rbac.LoginUrl, rbac.LoginResponse)
	if TestClient == nil {
		return
	}
	pageKeys := httptest.Responses{
		{Key: "status", Value: http.StatusOK},
		{Key: "message", Value: response.ResponseOkMessage},
		{Key: "data", Value: httptest.Responses{
			{Key: "pageSize", Value: 10},
			{Key: "page", Value: 1},
			{Key: "list", Value: []httptest.Responses{
				{
					{Key: "id", Value: web.TenancyAuthorityId, Type: "ge"},
					{Key: "authorityName", Value: "商户管理员"},
					{Key: "authorityType", Value: multi.TenancyAuthority},
					{Key: "parentId", Value: 0},
					{Key: "defaultRouter", Value: "dashboard"},
					{Key: "updatedAt", Value: "", Type: "notempty"},
					{Key: "createdAt", Value: "", Type: "notempty"},
				},
			}},
			{Key: "total", Value: 0, Type: "ge"},
		}},
	}
	requestParams := map[string]interface{}{"page": 1, "pageSize": 10, "orderBy": "id"}
	TestClient.GET(fmt.Sprintf("%s/getTenancyAuthorityList", url), pageKeys, requestParams)
}

func TestGetGeneralAuthorityList(t *testing.T) {
	if TestServer == nil {
		t.Error("测试服务初始化失败")
		return
	}

	TestClient = TestServer.GetTestLogin(t, rbac.LoginUrl, rbac.LoginResponse)
	if TestClient == nil {
		return
	}
	pageKeys := httptest.Responses{
		{Key: "status", Value: http.StatusOK},
		{Key: "message", Value: response.ResponseOkMessage},
		{Key: "data", Value: httptest.Responses{
			{Key: "pageSize", Value: 10},
			{Key: "page", Value: 1},
			{Key: "list", Value: []httptest.Responses{
				{
					{Key: "id", Value: web.DeviceAuthorityId, Type: "ge"},
					{Key: "authorityName", Value: "设备用户"},
					{Key: "authorityType", Value: multi.GeneralAuthority},
					{Key: "parentId", Value: 0},
					{Key: "defaultRouter", Value: "dashboard"},
					{Key: "updatedAt", Value: "", Type: "notempty"},
					{Key: "createdAt", Value: "", Type: "notempty"},
				},
				{
					{Key: "id", Value: web.LiteAuthorityId, Type: "ge"},
					{Key: "authorityName", Value: "小程序用户"},
					{Key: "authorityType", Value: multi.GeneralAuthority},
					{Key: "parentId", Value: 0},
					{Key: "defaultRouter", Value: "dashboard"},
					{Key: "updatedAt", Value: "", Type: "notempty"},
					{Key: "createdAt", Value: "", Type: "notempty"},
				},
			}},
			{Key: "total", Value: 0, Type: "ge"},
		}},
	}
	requestParams := map[string]interface{}{"page": 1, "pageSize": 10, "orderBy": "id"}
	TestClient.GET(fmt.Sprintf("%s/getGeneralAuthorityList", url), pageKeys, requestParams)
}

func TestCreate(t *testing.T) {
	if TestServer == nil {
		t.Error("测试服务初始化失败")
		return
	}

	TestClient = TestServer.GetTestLogin(t, rbac.LoginUrl, rbac.LoginResponse)
	if TestClient == nil {
		return
	}
	data := map[string]interface{}{
		"authorityName": "test_authorityName_for_create",
		"parentId":      0,
		"authorityType": multi.AdminAuthority,
	}
	id := Create(TestClient, data)
	if id == 0 {
		t.Fatalf("测试添加用户失败 id=%d", id)
	}
	defer Delete(TestClient, id)
}

func TestUpdate(t *testing.T) {
	if TestServer == nil {
		t.Error("测试服务初始化失败")
		return
	}

	TestClient = TestServer.GetTestLogin(t, rbac.LoginUrl, rbac.LoginResponse)
	if TestClient == nil {
		return
	}
	data := map[string]interface{}{
		"authorityName": "test_authorityName_for_update",
		"parentId":      0,
		"authorityType": multi.AdminAuthority,
	}
	id := Create(TestClient, data)
	if id == 0 {
		t.Fatalf("测试添加用户失败 id=%d", id)
	}
	defer Delete(TestClient, id)

	update := map[string]interface{}{
		"authorityName": "test_authorityName_for_update1",
		"parentId":      0,
		"authorityType": multi.AdminAuthority,
	}

	pageKeys := httptest.Responses{
		{Key: "status", Value: http.StatusOK},
		{Key: "message", Value: response.ResponseOkMessage},
	}
	TestClient.PUT(fmt.Sprintf("%s/updateAuthority/%d", url, id), pageKeys, update)
}

func TestCopyAuthority(t *testing.T) {
	if TestServer == nil {
		t.Error("测试服务初始化失败")
		return
	}

	TestClient = TestServer.GetTestLogin(t, rbac.LoginUrl, rbac.LoginResponse)
	if TestClient == nil {
		return
	}
	data := map[string]interface{}{
		"authorityName": "test_authorityName_for_copy",
		"parentId":      0,
		"authorityType": multi.AdminAuthority,
	}
	id := Create(TestClient, data)
	if id == 0 {
		t.Fatalf("测试添加用户失败 id=%d", id)
	}
	defer Delete(TestClient, id)

	pageKeys := httptest.Responses{
		{Key: "status", Value: http.StatusOK},
		{Key: "message", Value: response.ResponseOkMessage},
	}
	copy := map[string]interface{}{
		"authorityName": "test_authorityName_after_copy",
	}

	TestClient.POST(fmt.Sprintf("%s/copyAuthority/%d", url, id), pageKeys, copy)
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
	return TestClient.POST(fmt.Sprintf("%s/createAuthority", url), pageKeys, data).GetId()
}

func Delete(TestClient *httptest.Client, id uint) {
	pageKeys := httptest.Responses{
		{Key: "status", Value: http.StatusOK},
		{Key: "message", Value: response.ResponseOkMessage},
	}
	TestClient.DELETE(fmt.Sprintf("%s/deleteAuthority/%d", url, id), pageKeys)
}
