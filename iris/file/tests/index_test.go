package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/snowlyg/helper/tests"
	rbac "github.com/snowlyg/iris-admin-rbac/iris"
	"github.com/snowlyg/iris-admin-rbac/iris/file"
)

var (
	url = "/api/v1/file"
)

func TestUpload(t *testing.T) {
	if TestServer == nil {
		t.Errorf("测试服务初始化失败")
	}
	TestClient = TestServer.GetTestLogin(t, rbac.LoginUrl, nil)
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
	files := []tests.File{
		{
			Key:    "file",
			Path:   name,
			Reader: fh,
		},
	}
	local := file.GetPath(md5Name)
	pageKeys := tests.Responses{
		{Key: "code", Value: 2000},
		{Key: "message", Value: "请求成功"},
		{Key: "data", Value: tests.Responses{
			{Key: "local", Value: local},
			{Key: "qiniu", Value: ""},
		}},
	}

	TestClient.UPLOAD(url, pageKeys, files)

	local = filepath.Join(pwd, "static/upload")
	err = os.RemoveAll(local)
	if err != nil {
		t.Error(err)
	}
}
