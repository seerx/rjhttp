package rjh

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// Upload 上传对象
type Upload struct {
	Request *http.Request
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
	mFile, fh, err := u.Request.FormFile(field)
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

// FileInfo 文件信息
type FileInfo struct {
	FieldName string
	Files     []struct {
		Name string
		Size int64
	}
}

// GetFileInfo 获取上传文件的信息
func (u *Upload) GetFileInfo(field string) (*struct {
	Name string
	Size int64
}, error) {
	file, exists := u.Request.MultipartForm.File[field]
	if (!exists) || len(file) == 0 {
		return nil, fmt.Errorf("%s is not exists", field)
	}
	return &struct {
		Name string
		Size int64
	}{
		Name: file[0].Filename,
		Size: file[0].Size,
	}, nil
}

// GetFiles 获取文件列表
func (u *Upload) GetFiles() ([]*FileInfo, error) {
	var lst []*FileInfo

	for field, files := range u.Request.MultipartForm.File {
		info := &FileInfo{
			FieldName: field,
		}
		for _, file := range files {
			info.Files = append(info.Files, struct {
				Name string
				Size int64
			}{
				Name: file.Filename,
				Size: file.Size,
			})
		}
		if len(info.Files) > 0 {
			lst = append(lst, info)
		}
	}
	return lst, nil
}
