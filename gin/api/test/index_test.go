package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/httptest"
	rbac "github.com/snowlyg/iris-admin-rbac/gin"
	"github.com/snowlyg/iris-admin-rbac/gin/api"
	"github.com/snowlyg/iris-admin/server/database"
	"github.com/snowlyg/iris-admin/server/database/orm"
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
	if TestServer == nil {
		t.Error("测试服务初始化失败")
		return
	}

	TestClient = TestServer.GetTestLogin(t, rbac.LoginUrl, rbac.LoginResponse)
	if TestClient == nil {
		return
	}

	pageParams := getPageParams()
	routes, _ := TestServer.GetSources()
	for _, pageParam := range pageParams {
		t.Run(fmt.Sprintf("路由权限测试，第%d页", pageParam.Page), func(t *testing.T) {
			items, err := getPageApis(pageParam)
			if err != nil {
				t.Fatalf("获取路由权限错误")
			}
			pageKeys := httptest.Responses{
				{Key: "status", Value: http.StatusOK},
				{Key: "message", Value: pageParam.Message},
				{Key: "data", Value: httptest.Responses{
					{Key: "pageSize", Value: pageParam.PageSize},
					{Key: "page", Value: pageParam.Page},
					{Key: "list", Value: items},
					{Key: "total", Value: len(routes)},
				}},
			}
			requestParams := map[string]interface{}{"page": pageParam.Page, "pageSize": pageParam.PageSize}
			TestClient.GET(fmt.Sprintf("%s/getList", url), pageKeys, requestParams)
		})
	}
}

func TestGetAll(t *testing.T) {
	if TestServer == nil {
		t.Error("测试服务初始化失败")
		return
	}

	TestClient = TestServer.GetTestLogin(t, rbac.LoginUrl, rbac.LoginResponse)
	if TestClient == nil {
		return
	}

	routes, _ := TestServer.GetSources()
	t.Run("路由权限测试", func(t *testing.T) {
		items, err := getAllApis(multi.AdminAuthority)
		if err != nil {
			t.Fatalf("获取路由权限错误")
		}
		if len(routes) != len(items) {
			t.Errorf("接口需要返回 %d 个路由，实际返回了 %d 个数据", len(routes), len(items))
		}
		pageKeys := httptest.Responses{
			{Key: "status", Value: http.StatusOK},
			{Key: "message", Value: response.ResponseOkMessage},
			{Key: "data", Value: items},
		}
		requestParams := map[string]interface{}{"authorityType": multi.AdminAuthority}
		TestClient.GET(fmt.Sprintf("%s/getAll", url), pageKeys, requestParams)
	})
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
	if TestServer == nil {
		t.Error("测试服务初始化失败")
		return
	}

	TestClient = TestServer.GetTestLogin(t, rbac.LoginUrl, rbac.LoginResponse)
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

	pageKeys := httptest.Responses{
		{Key: "status", Value: http.StatusOK},
		{Key: "message", Value: response.ResponseOkMessage},
	}
	TestClient.PUT(fmt.Sprintf("%s/updateApi/%d", url, id), pageKeys, update)
}

func TestGetById(t *testing.T) {
	if TestServer == nil {
		t.Error("测试服务初始化失败")
		return
	}

	TestClient = TestServer.GetTestLogin(t, rbac.LoginUrl, rbac.LoginResponse)
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
		{Key: "status", Value: http.StatusOK},
		{Key: "message", Value: response.ResponseOkMessage},
		{Key: "data", Value: httptest.Responses{
			{Key: "id", Value: 1, Type: "ge"},
			{Key: "path", Value: data["path"].(string)},
			{Key: "apiGroup", Value: data["apiGroup"].(string)},
			{Key: "description", Value: data["description"].(string)},
			{Key: "method", Value: data["method"].(string)},
			{Key: "authorityType", Value: data["authorityType"].(int)},
			{Key: "updatedAt", Value: "", Type: "notempty"},
			{Key: "createdAt", Value: "", Type: "notempty"},
		},
		},
	}
	TestClient.GET(fmt.Sprintf("%s/getApiById/%d", url, id), pageKeys)
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
	return TestClient.POST(fmt.Sprintf("%s/createApi", url), pageKeys, data).GetId()
}

func Delete(TestClient *httptest.Client, id uint, data map[string]interface{}) {
	pageKeys := httptest.Responses{
		{Key: "status", Value: http.StatusOK},
		{Key: "message", Value: response.ResponseOkMessage},
	}
	TestClient.DELETE(fmt.Sprintf("%s/deleteApi/%d", url, id), pageKeys, data)
}

func getAllApis(authorityType int) ([]httptest.Responses, error) {
	routes := []httptest.Responses{}
	apis := api.PageResponse{}
	err := apis.Find(database.Instance(), api.AuthorityTypeScope(authorityType))
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
