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
	t.Run("test pagination", func(t *testing.T) {
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
					{Key: "perms", Value: []httptest.Responses{
						{
							{Key: "path", Value: "/api/v1/authority/getAuthorityList"},
							{Key: "method", Value: "GET"},
						},
						{
							{Key: "path", Value: "/api/v1/authority/getAdminAuthorityList"},
							{Key: "method", Value: "GET"},
						},
						{
							{Key: "path", Value: "/api/v1/authority/getTenancyAuthorityList"},
							{Key: "method", Value: "GET"},
						},
						{
							{Key: "path", Value: "/api/v1/authority/getGeneralAuthorityList"},
							{Key: "method", Value: "GET"},
						},
						{
							{Key: "path", Value: "/api/v1/api/getAll"},
							{Key: "method", Value: "GET"},
						},
						{
							{Key: "path", Value: "/api/v1/api/getApiById/:id"},
							{Key: "method", Value: "GET"},
						},
						{
							{Key: "path", Value: "/api/v1/api/getList"},
							{Key: "method", Value: "GET"},
						},
						{
							{Key: "path", Value: "/api/v1/admin/getAll"},
							{Key: "method", Value: "GET"},
						},
						{
							{Key: "path", Value: "/api/v1/admin/getAdmin/:id"},
							{Key: "method", Value: "GET"},
						},
						{
							{Key: "path", Value: "/api/v1/public/captcha"},
							{Key: "method", Value: "GET"},
						},
						{
							{Key: "path", Value: "/api/v1/public/clean"},
							{Key: "method", Value: "GET"},
						},
						{
							{Key: "path", Value: "/api/v1/public/logout"},
							{Key: "method", Value: "GET"},
						},
						{
							{Key: "path", Value: "/api/v1/profile"},
							{Key: "method", Value: "GET"},
						},
						{
							{Key: "path", Value: "/api/v1/oplog/getOplogList"},
							{Key: "method", Value: "GET"},
						},
						{
							{Key: "path", Value: "/api/v1/authority/createAuthority"},
							{Key: "method", Value: "POST"},
						},
						{
							{Key: "path", Value: "/api/v1/authority/copyAuthority/:id"},
							{Key: "method", Value: "POST"},
						},
						{
							{Key: "path", Value: "/api/v1/api/createApi"},
							{Key: "method", Value: "POST"},
						},
						{
							{Key: "path", Value: "/api/v1/admin/createAdmin"},
							{Key: "method", Value: "POST"},
						},
						{
							{Key: "path", Value: "/api/v1/profile/changeAvatar"},
							{Key: "method", Value: "POST"},
						},
						{
							{Key: "path", Value: "/api/v1/public/admin/login"},
							{Key: "method", Value: "POST"},
						},
						{
							{Key: "path", Value: "/api/v1/file/upload"},
							{Key: "method", Value: "POST"},
						},
						{
							{Key: "path", Value: "/api/v1/api/deleteApi/:id"},
							{Key: "method", Value: "DELETE"},
						},
						{
							{Key: "path", Value: "/api/v1/api/deleteApisByIds"},
							{Key: "method", Value: "DELETE"},
						},
						{
							{Key: "path", Value: "/api/v1/admin/deleteAdmin/:id"},
							{Key: "method", Value: "DELETE"},
						},
						{
							{Key: "path", Value: "/api/v1/authority/deleteAuthority/:id"},
							{Key: "method", Value: "DELETE"},
						},
						{
							{Key: "path", Value: "/api/v1/api/updateApi/:id"},
							{Key: "method", Value: "PUT"},
						},
						{
							{Key: "path", Value: "/api/v1/admin/updateAdmin/:id"},
							{Key: "method", Value: "PUT"},
						},
						{
							{Key: "path", Value: "/api/v1/authority/updateAuthority/:id"},
							{Key: "method", Value: "PUT"},
						},
					}},
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

	t.Run("test authorityName key", func(t *testing.T) {
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
		data := map[string]interface{}{"page": 1, "pageSize": 10, "orderBy": "id", "authorityName": "小程序用户"}
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
		"perms": [][]string{
			{"/api/v1/authority/updateAuthority/:id", "PUT"},
		},
	}
	id := Create(TestClient, data)
	if id == 0 {
		t.Fatalf("测试添加用户失败 id=%d", id)
	}
	defer Delete(TestClient, id)

	pageKeys := httptest.Responses{
		{Key: "pageSize", Value: 10},
		{Key: "page", Value: 1},
		{Key: "list", Value: []httptest.Responses{
			{
				{Key: "id", Value: 0, Type: "ge"},
				{Key: "uuid", Value: "test_authorityName_for_create"},
				{Key: "authorityName", Value: "test_authorityName_for_create"},
				{Key: "authorityType", Value: multi.AdminAuthority},
				{Key: "parentId", Value: 0},
				{Key: "defaultRouter", Value: "dashboard"},
				{Key: "updatedAt", Value: "", Type: "notempty"},
				{Key: "createdAt", Value: "", Type: "notempty"},
				{Key: "perms", Value: []httptest.Responses{
					{
						{Key: "path", Value: "/api/v1/authority/updateAuthority/:id"},
						{Key: "method", Value: "PUT"},
					},
				}},
			},
		}},
		{Key: "total", Value: 0, Type: "ge"},
	}

	dataPage := map[string]interface{}{"page": 1, "pageSize": 10, "orderBy": "id", "authorityName": "test_authorityName_for_create"}
	TestClient.GET(fmt.Sprintf("%s/getAuthorityList", url), httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, pageKeys), httptest.NewWithQueryObjectParamFunc(dataPage))

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
		"perms": [][]string{
			{"/api/v1/authority/updateAuthority/:id", "PUT"},
		},
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
		"perms": [][]string{
			{"/api/v1/authority/updateAuthority1/:id", "GET"},
		},
	}

	TestClient.PUT(fmt.Sprintf("%s/updateAuthority/%d", url, id), httptest.SuccessResponse, httptest.NewWithJsonParamFunc(requestParams))

	pageKeys := httptest.Responses{
		{Key: "pageSize", Value: 10},
		{Key: "page", Value: 1},
		{Key: "list", Value: []httptest.Responses{
			{
				{Key: "id", Value: 0, Type: "ge"},
				{Key: "uuid", Value: "test_authorityName_for_update"},
				{Key: "authorityName", Value: "test_authorityName_for_update1"},
				{Key: "authorityType", Value: multi.AdminAuthority},
				{Key: "parentId", Value: 0},
				{Key: "defaultRouter", Value: "dashboard"},
				{Key: "updatedAt", Value: "", Type: "notempty"},
				{Key: "createdAt", Value: "", Type: "notempty"},
				{Key: "perms", Value: []httptest.Responses{
					{
						{Key: "path", Value: "/api/v1/authority/updateAuthority1/:id"},
						{Key: "method", Value: "GET"},
					},
				}},
			},
		}},
		{Key: "total", Value: 0, Type: "ge"},
	}

	dataPage := map[string]interface{}{"page": 1, "pageSize": 10, "orderBy": "id", "authorityName": "test_authorityName_for_update"}
	TestClient.GET(fmt.Sprintf("%s/getAuthorityList", url), httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, pageKeys), httptest.NewWithQueryObjectParamFunc(dataPage))
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
