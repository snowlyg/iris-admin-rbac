package tests

import (
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/snowlyg/helper/str"
	"github.com/snowlyg/httptest"
	rbac "github.com/snowlyg/iris-admin-rbac/iris"
	"github.com/snowlyg/iris-admin-rbac/iris/file"
	"github.com/snowlyg/iris-admin/server/web"
	"github.com/snowlyg/iris-admin/server/web/web_gin/response"
)

var (
	url = "/api/v1/file"
)

func TestUpload(t *testing.T) {
	TestClient := httptest.Instance(t, TestServer.GetEngine(), str.Join("http://", web.CONFIG.System.Addr))
	TestClient.Login(rbac.LoginUrl, "", httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, rbac.LoginResponse))
	if TestClient == nil {
		return
	}

	name := "index_test.go"
	md5Name, err := file.GetFileName(name)
	if err != nil {
		t.Error(err)
		return
	}
	pwd, err := os.Getwd()
	if err != nil {
		t.Error(err)
		return
	}
	fullpath := filepath.Join(pwd, name)
	fh, err := os.Open(fullpath)
	if err != nil {
		t.Error(err)
		return
	}
	defer fh.Close()
	files := []httptest.File{
		{
			Key:    "file",
			Path:   name,
			Reader: fh,
		},
	}
	local := file.GetPath(md5Name)
	pageKeys := httptest.Responses{
		{Key: "local", Value: local},
		{Key: "qiniu", Value: ""},
	}

	TestClient.UPLOAD(url, httptest.NewResponses(http.StatusOK, response.ResponseOkMessage, pageKeys), httptest.NewWithFileParamFunc(files))

	local = filepath.Join(pwd, "static/upload")
	err = os.RemoveAll(local)
	if err != nil {
		t.Error(err)
	}
}
