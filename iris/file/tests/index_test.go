package tests

import (
	"os"
	"testing"

	"github.com/snowlyg/helper/tests"
	"github.com/snowlyg/iris-admin-rbac/iris/file"
)

var (
	loginUrl = "/api/v1/auth/login"
	url      = "/api/v1/file"
)

func TestUpload(t *testing.T) {
	if TestServer == nil {
		t.Errorf("TestServer is nil")
	}
	TestClient = TestServer.GetTestLogin(t, loginUrl, nil)
	if TestClient == nil {
		return
	}

	name := "mysqlPwd.txt"
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
