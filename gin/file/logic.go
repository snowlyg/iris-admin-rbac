package file

import (
	"mime/multipart"

	"github.com/snowlyg/iris-admin-rbac/gin/file/oss"
)

// uploadFile
func uploadFile(header *multipart.FileHeader) (string, error) {
	oss := oss.NewOss()
	filePath, _, err := oss.UploadFile(header)
	if err != nil {
		return "", err
	}

	return filePath, nil
}
