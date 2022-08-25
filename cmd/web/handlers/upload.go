package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

// UploadFile 上传文件，path 是存放上传文件基础路径
func UploadFile(r *http.Request, path string) (url string, err error) {
	file, header, err := r.FormFile("file")
	if err != nil {
		return
	}
	defer file.Close()

	// 获取文件后缀名
	suffix := filepath.Base(header.Filename)
	prefix := uuid.NewString()
	url = fmt.Sprintf("%s%s%s", path, prefix, suffix)

	f, err := os.OpenFile(url, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		return
	}

	return
}
