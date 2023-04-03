package oss

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/snowlyg/helper/str"
	"github.com/snowlyg/iris-admin-rbac/g"
	"github.com/snowlyg/iris-admin/server/zap_server"
)

type Local struct{}

func (l *Local) UploadFile(file *multipart.FileHeader) (string, string, error) {

	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = MD5V([]byte(name))
	// 拼接新文件名
	filename := str.Join(name, "_", time.Now().Local().Format("20060102150405"), ext)
	// 尝试创建此路径
	err := os.MkdirAll(g.RootPath, os.ModePerm)
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return "", "", err
	}
	// 拼接路径和文件名
	p := filepath.ToSlash(filepath.Join(g.RBAC_CONFIG.Path, filename))
	f, err := file.Open() // 读取文件
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return "", "", err
	}
	defer f.Close() // 创建文件 defer 关闭

	out, err := os.Create(filepath.ToSlash(filepath.Join(g.RootPath, filename)))
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())

		return "", "", err
	}
	defer out.Close() // 创建文件 defer 关闭

	_, err = io.Copy(out, f) // 传输（拷贝）文件
	if err != nil {
		zap_server.ZAPLOG.Error(err.Error())
		return "", "", err
	}
	return p, filename, nil
}

func (l *Local) DeleteFile(key string) error {
	p := filepath.ToSlash(filepath.Join(g.RootPath, key))
	if strings.Contains(p, g.RootPath) {
		if err := os.Remove(p); err != nil {
			zap_server.ZAPLOG.Error(err.Error())
			return err
		}
	}
	return nil
}

// MD5V md5加密
func MD5V(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}
