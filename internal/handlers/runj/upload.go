package runj

import (
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/seerx/rjhttp/pkg/rjh"
)

type Upload struct {
	request *http.Request
}

// StoreFile 存储文件
func (u *Upload) StoreFile(storePath string, field string) error {
	return u.StoreFileX(storePath, field, nil)
}

// StoreFileX 存储文件
func (u *Upload) StoreFileX(storePath string, field string, confirmFn func(*multipart.FileHeader) error) error {
	data, mfh, err := u.ReadFile(field)
	if err != nil {
		return err
	}
	if confirmFn != nil {
		if err := confirmFn(mfh); err != nil {
			// 不符合存储要求
			return err
		}
	}

	file, err := os.Create(storePath)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.Write(data); err != nil {
		return err
	}
	return nil
}

// ReadFile 读取文件内容
func (u *Upload) ReadFile(field string) ([]byte, *multipart.FileHeader, error) {
	mFile, fh, err := u.request.FormFile(field)
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		mFile.Close()
	}()
	data, err := ioutil.ReadAll(mFile)
	if err != nil {
		return nil, fh, err
	}
	return data, fh, nil
}

// injectUpload 上传辅助类注入函数
func injectUpload(arg map[string]interface{}) (*Upload, error) {
	request := rjh.ParseRequest(arg)
	//writer := ParseWriter(arg)
	//maxSize := ParseUploadMaxSize(arg)
	//request.Body = http.MaxBytesReader(writer, request.Body, maxSize)
	return &Upload{request: request}, nil
}
