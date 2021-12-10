package test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/snowlyg/helper/tests"
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
			items, err := getApis(pageParam)
			if err != nil {
				t.Fatalf("获取路由权限错误")
			}
			pageKeys := tests.Responses{
				{Key: "status", Value: http.StatusOK},
				{Key: "message", Value: pageParam.Message},
				{Key: "data", Value: tests.Responses{
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
		"name":        "update_test_route_name",
		"displayName": "测试描述信息",
		"description": "测试描述信息",
		"act":         "GET",
	}
	id := Create(client, data)
	if id == 0 {
		t.Fatalf("测试添加用户失败 id=%d", id)
	}
	defer Delete(client, id)

	update := map[string]interface{}{
		"name":        "update_test_route_name",
		"displayName": "更新测试描述信息",
		"description": "更新测试描述信息",
		"act":         "POST",
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
		"name":        "getbyid_test_route_name",
		"displayName": "测试描述信息",
		"description": "测试描述信息",
		"act":         "GET",
	}
	id := Create(client, data)
	if id == 0 {
		t.Fatalf("测试添加用户失败 id=%d", id)
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
			{Key: "act", Value: data["act"].(string)},
			{Key: "updatedAt", Value: "", Type: "notempty"},
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

func getApis(pageParam PageParam) ([]tests.Responses, error) {
	l := pageParam.PageLen
	routes := make([]tests.Responses, 0, l)
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
		api := tests.Responses{
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
