package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/helper/str"
	"github.com/snowlyg/httptest"
	rbac "github.com/snowlyg/iris-admin-rbac/iris"
	"github.com/snowlyg/iris-admin-rbac/iris/perm"
	"github.com/snowlyg/iris-admin/server/database"
	"github.com/snowlyg/iris-admin/server/database/orm"
	"github.com/snowlyg/iris-admin/server/web"
	"github.com/snowlyg/iris-admin/server/web/web_gin/response"
)

var (
	url = "/api/v1/perms"
)

type PageParam struct {
	Message  string
	Code     int
	PageSize int
	Page     int
	PageLen  int
}

func TestList(t *testing.T) {

	pageParams := getPageParams()
	routes, _ := TestServer.GetSources()
	for _, pageParam := range pageParams {
		t.Run(fmt.Sprintf("路由权限测试，第%d页", pageParam.Page), func(t *testing.T) {

			TestClient := httptest.Instance(t, TestServer.GetEngine(), str.Join("http://", web.CONFIG.System.Addr))
			TestClient.Login(rbac.LoginUrl, "", httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, rbac.LoginResponse))
			if TestClient == nil {
				return
			}
			items, err := getPerms(pageParam)
			if err != nil {
				t.Fatalf("获取路由权限错误")
			}
			pageKeys := httptest.Responses{
				{Key: "pageSize", Value: pageParam.PageSize},
				{Key: "page", Value: pageParam.Page},
				{Key: "items", Value: items},
				{Key: "total", Value: len(routes)},
			}

			requestParams := map[string]interface{}{"page": pageParam.Page, "pageSize": pageParam.PageSize}
			TestClient.GET(url, httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, pageKeys), httptest.NewWithQueryObjectParamFunc(requestParams))
		})
	}

}

func TestCreate(t *testing.T) {
	TestClient := httptest.Instance(t, TestServer.GetEngine(), str.Join("http://", web.CONFIG.System.Addr))
	TestClient.Login(rbac.LoginUrl, "", httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, rbac.LoginResponse))
	if TestClient == nil {
		return
	}

	data := map[string]interface{}{
		"name":        "test_route_name",
		"displayName": "测试描述信息",
		"description": "测试描述信息",
		"act":         "GET",
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
		"name":        "update_test_route_name",
		"displayName": "测试描述信息",
		"description": "测试描述信息",
		"act":         "GET",
	}
	id := Create(TestClient, data)
	if id == 0 {
		t.Fatalf("测试添加用户失败 id=%d", id)
	}
	defer Delete(TestClient, id)

	update := map[string]interface{}{
		"name":        "update_test_route_name",
		"displayName": "更新测试描述信息",
		"description": "更新测试描述信息",
		"act":         "POST",
	}

	TestClient.POST(fmt.Sprintf("%s/%d", url, id), httptest.SuccessResponse, httptest.NewWithJsonParamFunc(update))
}

func TestGetById(t *testing.T) {
	TestClient := httptest.Instance(t, TestServer.GetEngine(), str.Join("http://", web.CONFIG.System.Addr))
	TestClient.Login(rbac.LoginUrl, "", httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, rbac.LoginResponse))
	if TestClient == nil {
		return
	}

	data := map[string]interface{}{
		"name":        "getbyid_test_route_name",
		"displayName": "测试描述信息",
		"description": "测试描述信息",
		"act":         "GET",
	}
	id := Create(TestClient, data)
	if id == 0 {
		t.Fatalf("测试添加用户失败 id=%d", id)
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
			{Key: "act", Value: data["act"].(string)},
			{Key: "updatedAt", Value: "", Type: "notempty"},
			{Key: "createdAt", Value: "", Type: "notempty"},
		},
		},
	}
	TestClient.GET(fmt.Sprintf("%s/%d", url, id), pageKeys)
}

func Create(TestClient *httptest.Client, data map[string]interface{}) uint {
	pageKeys := httptest.IdKeys()
	TestClient.POST(url, httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, pageKeys), httptest.NewWithJsonParamFunc(data))
	return pageKeys.GetId()
}

func Delete(TestClient *httptest.Client, id uint) {
	TestClient.DELETE(fmt.Sprintf("%s/%d", url, id), httptest.SuccessResponse)
}

func getPerms(pageParam PageParam) ([]httptest.Responses, error) {
	l := pageParam.PageLen
	routes := make([]httptest.Responses, 0, l)
	req := &orm.Paginate{
		Page:     pageParam.Page,
		PageSize: pageParam.PageSize,
	}
	perms := &perm.PageResponse{}
	_, err := perms.Paginate(database.Instance(), req.PaginateScope())
	if err != nil {
		return routes, err
	}
	for _, route := range perms.Item {
		perm := httptest.Responses{
			{Key: "id", Value: route.Id},
			{Key: "name", Value: route.Name},
			{Key: "displayName", Value: route.DisplayName},
			{Key: "description", Value: route.Description},
			{Key: "act", Value: route.Act},
			{Key: "updatedAt", Value: route.UpdatedAt},
			{Key: "createdAt", Value: route.CreatedAt},
		}
		routes = append(routes, perm)
		l--
		if l == 0 {
			break
		}
	}

	return routes, err
}

func getPageParams() []PageParam {
	routes, _ := TestServer.GetSources()
	pageSize := 10
	size := len(routes) / pageSize
	other := len(routes) % pageSize
	if other > 0 {
		size++
	}
	pageParams := make([]PageParam, 0, size)
	for i := 1; i <= size; i++ {
		pageLen := pageSize
		if other > 0 && i == size {
			pageLen = other
		}
		pageParam := PageParam{
			Message:  response.ResponseOkMessage,
			Code:     http.StatusOK,
			PageSize: pageSize,
			PageLen:  pageLen,
			Page:     i,
		}
		pageParams = append(pageParams, pageParam)
	}
	return pageParams
}
