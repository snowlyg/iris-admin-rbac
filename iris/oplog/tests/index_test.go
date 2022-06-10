package tests

import (
	"net/http"
	"testing"

	"github.com/snowlyg/helper/str"
	"github.com/snowlyg/httptest"
	rbac "github.com/snowlyg/iris-admin-rbac/iris"
	"github.com/snowlyg/iris-admin/server/web"
	"github.com/snowlyg/iris-admin/server/web/web_gin/response"
)

var (
	url = "/api/v1/oplog"
)

type PageParam struct {
	Message  string
	Code     int
	PageSize int
	Page     int
	PageLen  int
}

func TestList(t *testing.T) {

	TestClient := httptest.Instance(t, TestServer.GetEngine(), str.Join("http://", web.CONFIG.System.Addr))
	TestClient.Login(rbac.LoginUrl, "", httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, rbac.LoginResponse))
	if TestClient == nil {
		return
	}
	pageKeys := httptest.Responses{
		{Key: "pageSize", Value: 10},
		{Key: "page", Value: 1},
		{Key: "items", Value: []httptest.Responses{}},
		{Key: "total", Value: 0, Type: "ge"},
	}
	TestClient.GET(url, httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, pageKeys), httptest.GetRequestFunc)
}
