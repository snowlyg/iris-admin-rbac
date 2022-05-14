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
)

var (
	url = "/api/v1/oplog" // url
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
			{Key: "list", Value: []httptest.Responses{}},
			{Key: "total", Value: 0, Type: "ge"},
		}},
	}
	requestParams := map[string]interface{}{"page": 1, "pageSize": 10, "orderBy": "id"}
	TestClient.GET(fmt.Sprintf("%s/getOplogList", url), pageKeys, httptest.NewWithQueryObjectParamFunc(requestParams))
}
