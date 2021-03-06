package test

import (
	"fmt"
	"net/http"
	"path/filepath"
	"testing"

	"github.com/snowlyg/helper/str"
	"github.com/snowlyg/httptest"
	rbac "github.com/snowlyg/iris-admin-rbac/gin"
	"github.com/snowlyg/iris-admin/g"
	"github.com/snowlyg/iris-admin/server/web"
	"github.com/snowlyg/iris-admin/server/web/web_gin/response"
)

var (
	url = "/api/v1/admin"
)

func TestList(t *testing.T) {
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
				{Key: "id", Value: 1},
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
				{Key: "deletedAt", Value: ""},
			},
		}},
		{Key: "total", Value: 1},
	}
	TestClient.GET(fmt.Sprintf("%s/getAll", url), httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, pageKeys), httptest.GetRequestFunc)
}

func TestCreate(t *testing.T) {
	TestClient := httptest.Instance(t, TestServer.GetEngine(), str.Join("http://", web.CONFIG.System.Addr))
	TestClient.Login(rbac.LoginUrl, "", httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, rbac.LoginResponse))
	if TestClient == nil {
		return
	}

	data := map[string]interface{}{
		"nickName":     "测试名称",
		"username":     "create_test_username",
		"authorityIds": []string{"super_admin"},
		"email":        "get@admin.com",
		"phone":        "13800138001",
		"password":     "123456",
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
		"nickName":     "测试名称",
		"username":     "create_test_username_for_update",
		"authorityIds": []string{"super_admin"},
		"email":        "get@admin.com",
		"phone":        "13800138001",
		"password":     "123456",
	}
	id := Create(TestClient, data)
	if id == 0 {
		t.Fatalf("测试添加用户失败 id=%d", id)
	}
	defer Delete(TestClient, id)

	update := map[string]interface{}{
		"nickName": "测试名称",
		"email":    "get@admin.com",
		"phone":    "13800138003",
		"password": "123456",
	}

	TestClient.PUT(fmt.Sprintf("%s/updateAdmin/%d", url, id), httptest.SuccessResponse, httptest.NewWithJsonParamFunc(update))
}

func TestGetById(t *testing.T) {

	TestClient := httptest.Instance(t, TestServer.GetEngine(), str.Join("http://", web.CONFIG.System.Addr))
	TestClient.Login(rbac.LoginUrl, "", httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, rbac.LoginResponse))
	if TestClient == nil {
		return
	}
	data := map[string]interface{}{
		"nickName":     "测试名称",
		"username":     "create_test_username_for_get",
		"email":        "get@admin.com",
		"phone":        "13800138001",
		"authorityIds": []string{"super_admin"},
		"password":     "123456",
	}
	id := Create(TestClient, data)
	if id == 0 {
		t.Fatalf("测试添加用户失败 id=%d", id)
	}
	defer Delete(TestClient, id)
	pageKeys := httptest.Responses{
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
	}
	TestClient.GET(fmt.Sprintf("%s/getAdmin/%d", url, id), httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, pageKeys))
}

func TestChangeAvatar(t *testing.T) {

	TestClient := httptest.Instance(t, TestServer.GetEngine(), str.Join("http://", web.CONFIG.System.Addr))
	TestClient.Login(rbac.LoginUrl, "", httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, rbac.LoginResponse))
	if TestClient == nil {
		return
	}
	data := map[string]interface{}{
		"headerImg": "/avatar.png",
	}
	TestClient.POST(fmt.Sprintf("%s/changeAvatar", url), httptest.SuccessResponse, httptest.NewWithJsonParamFunc(data))

	profile := httptest.Responses{
		{Key: "id", Value: 1, Type: "ge"},
		{Key: "nickName", Value: "超级管理员"},
		{Key: "username", Value: "admin"},
		{Key: "headerImg", Value: filepath.ToSlash(web.ToStaticUrl("/avatar.png"))},
		{Key: "status", Value: g.StatusTrue},
		{Key: "isShow", Value: g.StatusFalse},
		{Key: "phone", Value: "13800138000"},
		{Key: "email", Value: "admin@admin.com"},
		{Key: "authorities", Value: []string{"超级管理员"}},
		{Key: "updatedAt", Value: "", Type: "notempty"},
		{Key: "createdAt", Value: "", Type: "notempty"},
	}
	TestClient.GET(fmt.Sprintf("%s/profile", url), httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, profile))
}

func Create(TestClient *httptest.Client, data map[string]interface{}) uint {
	pageKeys := httptest.IdKeys()
	TestClient.POST(fmt.Sprintf("%s/createAdmin", url), httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, pageKeys), httptest.NewWithJsonParamFunc(data))
	return pageKeys.GetId()
}

func Delete(TestClient *httptest.Client, id uint) {
	TestClient.DELETE(fmt.Sprintf("%s/deleteAdmin/%d", url, id), httptest.SuccessResponse)
}
