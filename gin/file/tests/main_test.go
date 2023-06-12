package tests

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/snowlyg/httptest"
	"github.com/snowlyg/iris-admin-rbac/g"
	rbac "github.com/snowlyg/iris-admin-rbac/gin"
	"github.com/snowlyg/iris-admin/server/cache"
	"github.com/snowlyg/iris-admin/server/web"
	"github.com/snowlyg/iris-admin/server/web/common"
	"github.com/snowlyg/iris-admin/server/web/web_gin"
)

var TestServer *web_gin.WebServer

func TestMain(m *testing.M) {
	os.Setenv("driverType", "redis")
	cache.CONFIG.DB = 5
	redisPwd := os.Getenv("redisProPwd")
	if redisPwd != "" {
		cache.CONFIG.Password = redisPwd
	}
	redisAddr := os.Getenv("redisAddr")
	if redisAddr != "" {
		cache.CONFIG.Addr = os.Getenv("redisAddr")
	}
	cache.Recover()
	gopath := os.Getenv("GOPATH")
	g.RootPath = filepath.ToSlash(filepath.Join(gopath, "/src/github.com/snowlyg/iris-admin-rbac", strings.TrimSpace(g.RBAC_CONFIG.Path)))

	web.CONFIG.Menu.Uri = "/api/v1/authority/getAuthorityList"
	web.CONFIG.Menu.Method = "GET"
	web.Recover()
	var uuid string
	uuid, TestServer = common.BeforeTestMainGin(rbac.PartyFunc, rbac.SeedFunc)
	code := m.Run()
	common.AfterTestMain(uuid, true)
	cache.Remove()
	os.Exit(code)
}

func UploadMedia(auth *httptest.Client, name string, status int, message string) string {
	fh, _ := os.Open("./" + name)
	defer fh.Close()
	uf := []httptest.File{{Key: "file", Path: name, Reader: fh}}
	url := "/api/v1/file/upload"

	src := httptest.Responses{
		{Key: "src", Value: "", Type: "notempty"},
	}
	pageKeys := httptest.NewResponses(status, message, src)
	auth.UPLOAD(url, pageKeys, httptest.NewWithFileParamFunc(uf))
	return src.GetString("src")
}
