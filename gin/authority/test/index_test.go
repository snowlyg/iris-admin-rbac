package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/helper/str"
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
	t.Run("test pagination",func(t *testing.T) {
		TestClient := httptest.Instance(t, TestServer.GetEngine(), str.Join("http://", web.CONFIG.System.Addr))
		TestClient.Login(rbac.LoginUrl, "", httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, rbac.LoginResponse))
		if TestClient == nil {
			return
		}
		pageKeys := httptest.Responses{
			{Key: "pageSize", Value: 10},
			{Key: "page", Value: 1},
			{Key: "list", Value: []httptest.Responses{
				{
					{Key: "id", Value: 0, Type: "ge"},
					{Key: "uuid", Value: "device_admin"},
					{Key: "authorityName", Value: "设备用户"},
					{Key: "authorityType", Value: multi.GeneralAuthority},
					{Key: "parentId", Value: 0},
					{Key: "defaultRouter", Value: "dashboard"},
					{Key: "updatedAt", Value: "", Type: "notempty"},
					{Key: "createdAt", Value: "", Type: "notempty"},
				},
				{
					{Key: "id", Value: 0, Type: "ge"},
					{Key: "uuid", Value: "mini_admin"},
					{Key: "authorityName", Value: "小程序用户"},
					{Key: "authorityType", Value: multi.GeneralAuthority},
					{Key: "parentId", Value: 0},
					{Key: "defaultRouter", Value: "dashboard"},
					{Key: "updatedAt", Value: "", Type: "notempty"},
					{Key: "createdAt", Value: "", Type: "notempty"},
				},
				{
					{Key: "id", Value: 0, Type: "ge"},
					{Key: "uuid", Value: "tenancy_admin"},
					{Key: "authorityName", Value: "商户管理员"},
					{Key: "authorityType", Value: multi.TenancyAuthority},
					{Key: "parentId", Value: 0},
					{Key: "defaultRouter", Value: "dashboard"},
					{Key: "updatedAt", Value: "", Type: "notempty"},
					{Key: "createdAt", Value: "", Type: "notempty"},
				},
				{
					{Key: "id", Value: 0, Type: "ge"},
					{Key: "uuid", Value: "super_admin"},
					{Key: "authorityName", Value: "超级管理员"},
					{Key: "authorityType", Value: multi.AdminAuthority},
					{Key: "parentId", Value: 0},
					{Key: "defaultRouter", Value: "dashboard"},
					{Key: "updatedAt", Value: "", Type: "notempty"},
					{Key: "createdAt", Value: "", Type: "notempty"},
				},
			}},
			{Key: "total", Value: 0, Type: "ge"},
		}
		data := map[string]interface{}{"page": 1, "pageSize": 10, "orderBy": "id"}
		TestClient.GET(fmt.Sprintf("%s/getAuthorityList", url), httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, pageKeys), httptest.NewWithQueryObjectParamFunc(data))
	})

	t.Run("test authorityName key",func(t *testing.T) {
				TestClient := httptest.Instance(t, TestServer.GetEngine(), str.Join("http://", web.CONFIG.System.Addr))
		TestClient.Login(rbac.LoginUrl, "", httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, rbac.LoginResponse))
		if TestClient == nil {
			return
		}
		pageKeys := httptest.Responses{
			{Key: "pageSize", Value: 10},
			{Key: "page", Value: 1},
			{Key: "list", Value: []httptest.Responses{
				{
					{Key: "id", Value: 0, Type: "ge"},
					{Key: "uuid", Value: "mini_admin"},
					{Key: "authorityName", Value: "小程序用户"},
					{Key: "authorityType", Value: multi.GeneralAuthority},
					{Key: "parentId", Value: 0},
					{Key: "defaultRouter", Value: "dashboard"},
					{Key: "updatedAt", Value: "", Type: "notempty"},
					{Key: "createdAt", Value: "", Type: "notempty"},
				},
			}},
			{Key: "total", Value: 0, Type: "ge"},
		}
		data := map[string]interface{}{"page": 1, "pageSize": 10, "orderBy": "id","authorityName":"小程序用户"}
		TestClient.GET(fmt.Sprintf("%s/getAuthorityList", url), httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, pageKeys), httptest.NewWithQueryObjectParamFunc(data))
	})
}

func TestGetAdminAuthorityList(t *testing.T) {

	TestClient := httptest.Instance(t, TestServer.GetEngine(), str.Join("http://", web.CONFIG.System.Addr))
	TestClient.Login(rbac.LoginUrl, "", httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, rbac.LoginResponse))
	if TestClient == nil {
		return
	}
	pageKeys := httptest.Responses{
		{Key: "pageSize", Value: 10},
		{Key: "page", Value: 1},
		{Key: "list", Value: []httptest.Responses{
			{
				{Key: "id", Value: 0, Type: "ge"},
				{Key: "uuid", Value: "super_admin"},
				{Key: "authorityName", Value: "超级管理员"},
				{Key: "authorityType", Value: multi.AdminAuthority},
				{Key: "parentId", Value: 0},
				{Key: "defaultRouter", Value: "dashboard"},
				{Key: "updatedAt", Value: "", Type: "notempty"},
				{Key: "createdAt", Value: "", Type: "notempty"},
			},
		}},
		{Key: "total", Value: 0, Type: "ge"},
	}
	data := map[string]interface{}{"page": 1, "pageSize": 10, "orderBy": "id"}
	TestClient.GET(fmt.Sprintf("%s/getAdminAuthorityList", url), httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, pageKeys), httptest.NewWithQueryObjectParamFunc(data))
}

func TestGetTenancyAuthorityList(t *testing.T) {

	TestClient := httptest.Instance(t, TestServer.GetEngine(), str.Join("http://", web.CONFIG.System.Addr))
	TestClient.Login(rbac.LoginUrl, "", httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, rbac.LoginResponse))
	if TestClient == nil {
		return
	}
	pageKeys := httptest.Responses{
		{Key: "pageSize", Value: 10},
		{Key: "page", Value: 1},
		{Key: "list", Value: []httptest.Responses{
			{
				{Key: "id", Value: 0, Type: "ge"},
				{Key: "uuid", Value: "tenancy_admin"},
				{Key: "authorityName", Value: "商户管理员"},
				{Key: "authorityType", Value: multi.TenancyAuthority},
				{Key: "parentId", Value: 0},
				{Key: "defaultRouter", Value: "dashboard"},
				{Key: "updatedAt", Value: "", Type: "notempty"},
				{Key: "createdAt", Value: "", Type: "notempty"},
			},
		}},
		{Key: "total", Value: 0, Type: "ge"},
	}
	data := map[string]interface{}{"page": 1, "pageSize": 10, "orderBy": "id"}
	TestClient.GET(fmt.Sprintf("%s/getTenancyAuthorityList", url), httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, pageKeys), httptest.NewWithQueryObjectParamFunc(data))
}

func TestGetGeneralAuthorityList(t *testing.T) {

	TestClient := httptest.Instance(t, TestServer.GetEngine(), str.Join("http://", web.CONFIG.System.Addr))
	TestClient.Login(rbac.LoginUrl, "", httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, rbac.LoginResponse))
	if TestClient == nil {
		return
	}
	pageKeys := httptest.Responses{
		{Key: "pageSize", Value: 10},
		{Key: "page", Value: 1},
		{Key: "list", Value: []httptest.Responses{
			{
				{Key: "id", Value: 0, Type: "ge"},
				{Key: "uuid", Value: "device_admin"},
				{Key: "authorityName", Value: "设备用户"},
				{Key: "authorityType", Value: multi.GeneralAuthority},
				{Key: "parentId", Value: 0},
				{Key: "defaultRouter", Value: "dashboard"},
				{Key: "updatedAt", Value: "", Type: "notempty"},
				{Key: "createdAt", Value: "", Type: "notempty"},
			},
			{
				{Key: "id", Value: 0, Type: "ge"},
				{Key: "uuid", Value: "mini_admin"},
				{Key: "authorityName", Value: "小程序用户"},
				{Key: "authorityType", Value: multi.GeneralAuthority},
				{Key: "parentId", Value: 0},
				{Key: "defaultRouter", Value: "dashboard"},
				{Key: "updatedAt", Value: "", Type: "notempty"},
				{Key: "createdAt", Value: "", Type: "notempty"},
			},
		}},
		{Key: "total", Value: 0, Type: "ge"},
	}
	data := map[string]interface{}{"page": 1, "pageSize": 10, "orderBy": "id"}
	TestClient.GET(fmt.Sprintf("%s/getGeneralAuthorityList", url), httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, pageKeys), httptest.NewWithQueryObjectParamFunc(data))
}

func TestCreate(t *testing.T) {

	TestClient := httptest.Instance(t, TestServer.GetEngine(), str.Join("http://", web.CONFIG.System.Addr))
	TestClient.Login(rbac.LoginUrl, "", httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, rbac.LoginResponse))
	if TestClient == nil {
		return
	}
	data := map[string]interface{}{
		"uuid":          "test_authorityName_for_create",
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
	TestClient := httptest.Instance(t, TestServer.GetEngine(), str.Join("http://", web.CONFIG.System.Addr))
	TestClient.Login(rbac.LoginUrl, "", httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, rbac.LoginResponse))
	if TestClient == nil {
		return
	}
	data := map[string]interface{}{
		"uuid":          "test_authorityName_for_update",
		"authorityName": "test_authorityName_for_update",
		"parentId":      0,
		"authorityType": multi.AdminAuthority,
	}
	id := Create(TestClient, data)
	if id == 0 {
		t.Fatalf("测试添加用户失败 id=%d", id)
	}
	defer Delete(TestClient, id)

	requestParams := map[string]interface{}{
		"authorityName": "test_authorityName_for_update1",
		"parentId":      0,
		"authorityType": multi.AdminAuthority,
	}

	TestClient.PUT(fmt.Sprintf("%s/updateAuthority/%d", url, id), httptest.SuccessResponse, httptest.NewWithJsonParamFunc(requestParams))
}

func TestCopyAuthority(t *testing.T) {

	TestClient := httptest.Instance(t, TestServer.GetEngine(), str.Join("http://", web.CONFIG.System.Addr))
	TestClient.Login(rbac.LoginUrl, "", httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, rbac.LoginResponse))
	if TestClient == nil {
		return
	}
	data := map[string]interface{}{
		"uuid":          "test_authorityName_for_copy",
		"authorityName": "test_authorityName_for_copy",
		"parentId":      0,
		"authorityType": multi.AdminAuthority,
	}
	id := Create(TestClient, data)
	if id == 0 {
		t.Fatalf("测试添加用户失败 id=%d", id)
	}
	defer Delete(TestClient, id)

	update := map[string]interface{}{
		"uuid":          "test_authorityName_after_copy",
		"authorityName": "test_authorityName_after_copy",
	}

	TestClient.POST(fmt.Sprintf("%s/copyAuthority/%d", url, id), httptest.SuccessResponse, httptest.NewWithJsonParamFunc(update))
}

func Create(TestClient *httptest.Client, data map[string]interface{}) uint {
	pageKeys := httptest.IdKeys()
	TestClient.POST(fmt.Sprintf("%s/createAuthority", url), httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, pageKeys), httptest.NewWithJsonParamFunc(data))
	return pageKeys.GetId()
}

func Delete(TestClient *httptest.Client, id uint) {

	TestClient.DELETE(fmt.Sprintf("%s/deleteAuthority/%d", url, id), httptest.SuccessResponse)
}
