package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/helper/str"
	"github.com/snowlyg/httptest"
	rbac "github.com/snowlyg/iris-admin-rbac/gin"
	"github.com/snowlyg/iris-admin-rbac/gin/api"
	"github.com/snowlyg/iris-admin/server/database"
	"github.com/snowlyg/iris-admin/server/database/orm"
	"github.com/snowlyg/iris-admin/server/web"
	"github.com/snowlyg/iris-admin/server/web/web_gin/response"
	"github.com/snowlyg/multi"
)

var (
	url = "/api/v1/api"
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
			items, err := getPageApis(pageParam)
			if err != nil {
				t.Fatalf("获取路由权限错误")
			}
			pageKeys := httptest.Responses{
				{Key: "pageSize", Value: pageParam.PageSize},
				{Key: "page", Value: pageParam.Page},
				{Key: "list", Value: items},
				{Key: "total", Value: len(routes)},
			}
			data := map[string]interface{}{"page": pageParam.Page, "pageSize": pageParam.PageSize}
			TestClient.GET(fmt.Sprintf("%s/getList", url), httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, pageKeys), httptest.NewWithQueryObjectParamFunc(data))
		})
	}

	t.Run("路由权限测试，path key", func(t *testing.T) {
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
					{Key: "path", Value: "/api/v1/authority/getAuthorityList"},
					{Key: "description", Value: "GetAuthorityList"},
					{Key: "apiGroup", Value: "authority"},
					{Key: "method", Value: "GET"},
					{Key: "authorityType", Value: 1},
					{Key: "updatedAt", Value: "", Type: "notempty"},
					{Key: "createdAt", Value: "", Type: "notempty"}},
			}},
			{Key: "total", Value: 1},
		}
		data := map[string]interface{}{"page": 1, "pageSize": 10, "path": "/api/v1/authority/getAuthorityL"}
		TestClient.GET(fmt.Sprintf("%s/getList", url), httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, pageKeys), httptest.NewWithQueryObjectParamFunc(data))
	})

	t.Run("路由权限测试，sort key", func(t *testing.T) {
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
					{Key: "id", Value: 14},
					{Key: "path", Value: "/api/v1/oplog/getOplogList"},
					{Key: "description", Value: "GetOplogList"},
					{Key: "apiGroup", Value: "oplog"},
					{Key: "method", Value: "GET"},
					{Key: "authorityType", Value: 1},
					{Key: "updatedAt", Value: "", Type: "notempty"},
					{Key: "createdAt", Value: "", Type: "notempty"}},
			}, Length: 1},
			{Key: "total", Value: 14},
		}
		data := map[string]interface{}{"page": 1, "pageSize": 10, "method": "GET", "sort": "", "orderBy": "id"}
		TestClient.GET(fmt.Sprintf("%s/getList", url), httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, pageKeys), httptest.NewWithQueryObjectParamFunc(data))
	})

	t.Run("路由权限测试，method key", func(t *testing.T) {
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
					{Key: "path", Value: "/api/v1/authority/getAuthorityList"},
					{Key: "description", Value: "GetAuthorityList"},
					{Key: "apiGroup", Value: "authority"},
					{Key: "method", Value: "GET"},
					{Key: "authorityType", Value: 1},
					{Key: "updatedAt", Value: "", Type: "notempty"},
					{Key: "createdAt", Value: "", Type: "notempty"}},
			}, Length: 1},
			{Key: "total", Value: 14},
		}
		data := map[string]interface{}{"page": 1, "pageSize": 10, "method": "GET"}
		TestClient.GET(fmt.Sprintf("%s/getList", url), httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, pageKeys), httptest.NewWithQueryObjectParamFunc(data))
	})

	t.Run("路由权限测试，apiGroup key", func(t *testing.T) {
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
					{Key: "path", Value: "/api/v1/authority/getAuthorityList"},
					{Key: "description", Value: "GetAuthorityList"},
					{Key: "apiGroup", Value: "authority"},
					{Key: "method", Value: "GET"},
					{Key: "authorityType", Value: 1},
					{Key: "updatedAt", Value: "", Type: "notempty"},
					{Key: "createdAt", Value: "", Type: "notempty"}},
			}, Length: 1},
			{Key: "total", Value: 8},
		}
		data := map[string]interface{}{"page": 1, "pageSize": 10, "apiGroup": "authority"}
		TestClient.GET(fmt.Sprintf("%s/getList", url), httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, pageKeys), httptest.NewWithQueryObjectParamFunc(data))
	})
}

func TestGetAll(t *testing.T) {

	TestClient := httptest.Instance(t, TestServer.GetEngine(), str.Join("http://", web.CONFIG.System.Addr))
	TestClient.Login(rbac.LoginUrl, "", httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, rbac.LoginResponse))
	if TestClient == nil {
		return
	}

	routes, _ := TestServer.GetSources()
	t.Run("路由权限测试", func(t *testing.T) {
		items, err := getAllApis(multi.AdminAuthority)
		if err != nil {
			t.Fatalf("获取路由权限错误")
		}
		if 5 != len(items) {
			t.Errorf("接口需要返回 %d 个路由，实际返回了 %d 个数据", len(routes), len(items))
		}

		data := map[string]interface{}{"authorityType": multi.AdminAuthority}
		TestClient.GET(fmt.Sprintf("%s/getAll", url), httptest.NewResponsesWithLength(http.StatusOK, response.ResponseOkMessage, items, len(items)), httptest.NewWithQueryObjectParamFunc(data))
	})
}

func TestCreate(t *testing.T) {

	TestClient := httptest.Instance(t, TestServer.GetEngine(), str.Join("http://", web.CONFIG.System.Addr))
	TestClient.Login(rbac.LoginUrl, "", httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, rbac.LoginResponse))
	if TestClient == nil {
		return
	}
	data := map[string]interface{}{
		"path":          "getbyid_test_route_name",
		"apiGroup":      "apiGroup",
		"description":   "测试描述信息",
		"method":        "GET",
		"authorityType": multi.AdminAuthority,
	}
	id := Create(TestClient, data)
	if id == 0 {
		t.Fatalf("测试添加用户失败 id=%d", id)
	}
	defer Delete(TestClient, id, data)
}

func TestUpdate(t *testing.T) {

	TestClient := httptest.Instance(t, TestServer.GetEngine(), str.Join("http://", web.CONFIG.System.Addr))
	TestClient.Login(rbac.LoginUrl, "", httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, rbac.LoginResponse))
	if TestClient == nil {
		return
	}
	data := map[string]interface{}{
		"path":          "getbyid_test_route_name",
		"apiGroup":      "apiGroup",
		"description":   "测试描述信息",
		"method":        "GET",
		"authorityType": multi.AdminAuthority,
	}
	id := Create(TestClient, data)
	if id == 0 {
		t.Fatalf("测试添加用户失败 id=%d", id)
	}
	defer Delete(TestClient, id, data)

	update := map[string]interface{}{
		"path":          "getbyid_test_route_name",
		"apiGroup":      "apiGroup",
		"description":   "测试描述信息",
		"method":        "GET",
		"authorityType": multi.AdminAuthority,
	}

	TestClient.PUT(fmt.Sprintf("%s/updateApi/%d", url, id), httptest.SuccessResponse, httptest.NewWithJsonParamFunc(update))
}

func TestGetById(t *testing.T) {

	TestClient := httptest.Instance(t, TestServer.GetEngine(), str.Join("http://", web.CONFIG.System.Addr))
	TestClient.Login(rbac.LoginUrl, "", httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, rbac.LoginResponse))
	if TestClient == nil {
		return
	}
	data := map[string]interface{}{
		"path":          "getbyid_test_route_name",
		"apiGroup":      "apiGroup",
		"description":   "测试描述信息",
		"method":        "GET",
		"authorityType": multi.AdminAuthority,
	}
	id := Create(TestClient, data)
	if id == 0 {
		t.Fatalf("测试添加用户失败 id=%d", id)
	}
	defer Delete(TestClient, id, data)

	pageKeys := httptest.Responses{
		{Key: "id", Value: 1, Type: "ge"},
		{Key: "path", Value: data["path"].(string)},
		{Key: "apiGroup", Value: data["apiGroup"].(string)},
		{Key: "description", Value: data["description"].(string)},
		{Key: "method", Value: data["method"].(string)},
		{Key: "authorityType", Value: data["authorityType"].(int)},
		{Key: "updatedAt", Value: "", Type: "notempty"},
		{Key: "createdAt", Value: "", Type: "notempty"},
	}
	TestClient.GET(fmt.Sprintf("%s/getApiById/%d", url, id), httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, pageKeys))
}

func Create(TestClient *httptest.Client, data map[string]interface{}) uint {
	pageKeys := httptest.IdKeys()
	TestClient.POST(fmt.Sprintf("%s/createApi", url), httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, pageKeys), httptest.NewWithJsonParamFunc(data))
	return pageKeys.GetId()
}

func Delete(TestClient *httptest.Client, id uint, data map[string]interface{}) {
	TestClient.DELETE(fmt.Sprintf("%s/deleteApi/%d", url, id), httptest.SuccessResponse, httptest.NewWithJsonParamFunc(data))
}

func getAllApis(authorityType int) ([]httptest.Responses, error) {
	routes := []httptest.Responses{}
	apis := api.PageResponse{}
	err := apis.Find(database.Instance(), api.AuthorityTypeScope(authorityType))
	if err != nil {
		return routes, err
	}

	routers := api.FormatApis(apis.Item)

	for _, route := range routers {
		children := []httptest.Responses{}
		if len(route.Children) > 0 {
			for _, routeChild := range route.Children {
				child := httptest.Responses{
					{Key: "id", Value: routeChild.Id},
					{Key: "path", Value: routeChild.Path},
					{Key: "description", Value: routeChild.Description},
					{Key: "apiGroup", Value: routeChild.ApiGroup},
					{Key: "method", Value: routeChild.Method},
					{Key: "authorityType", Value: routeChild.AuthorityType},
					{Key: "updatedAt", Value: routeChild.UpdatedAt},
					{Key: "createdAt", Value: routeChild.CreatedAt},
				}
				children = append(children, child)
			}
		}
		api := httptest.Responses{
			{Key: "id", Value: route.Id},
			{Key: "path", Value: route.Path},
			{Key: "description", Value: route.Description},
			{Key: "apiGroup", Value: route.ApiGroup},
			{Key: "method", Value: route.Method},
			{Key: "authorityType", Value: route.AuthorityType},
			{Key: "updatedAt", Value: route.UpdatedAt},
			{Key: "createdAt", Value: route.CreatedAt},
			{Key: "children", Value: children},
		}

		routes = append(routes, api)
	}

	return routes, err
}

func getPageApis(pageParam PageParam) ([]httptest.Responses, error) {
	l := pageParam.PageLen
	routes := make([]httptest.Responses, 0, l)
	req := &orm.Paginate{
		Page:     pageParam.Page,
		PageSize: pageParam.PageSize,
	}
	apis := api.PageResponse{}
	_, err := apis.Paginate(database.Instance(), req.PaginateScope())
	if err != nil {
		return routes, err
	}

	for _, route := range apis.Item {
		api := httptest.Responses{
			{Key: "id", Value: route.Id},
			{Key: "path", Value: route.Path},
			{Key: "description", Value: route.Description},
			{Key: "apiGroup", Value: route.ApiGroup},
			{Key: "method", Value: route.Method},
			{Key: "authorityType", Value: route.AuthorityType},
			{Key: "updatedAt", Value: route.UpdatedAt},
			{Key: "createdAt", Value: route.CreatedAt},
		}
		routes = append(routes, api)
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
			Code:     2000,
			PageSize: pageSize,
			PageLen:  pageLen,
			Page:     i,
		}
		pageParams = append(pageParams, pageParam)
	}
	return pageParams
}
