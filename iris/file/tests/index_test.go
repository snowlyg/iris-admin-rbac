package tests

import (
	"os"
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
	fh, err := os.Open("D:/admin/go/src/github.com/snowlyg/iris-admin-rbac/iris/file/tests/" + name)
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
	pageKeys := tests.Responses{
		{Key: "code", Value: 2000},
		{Key: "message", Value: "请求成功"},
		{Key: "data", Value: tests.Responses{
			{Key: "local", Value: file.GetPath(md5Name)},
			{Key: "qiniu", Value: ""},
		}},
	}

	TestClient.UPLOAD(url, pageKeys, files)
}
