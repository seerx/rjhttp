package rjh

import (
	"net/http"
	"reflect"
)

// ResponseBinary 返回二进制文件(含图片)时，返回直结构
type ResponseBinary struct {
	Binary string `json:"binary" rj:"desc:Return binary data, for example: http.ResponseWriter.Write(...)"`
}

var binType = reflect.TypeOf((*ResponseBinary)(nil)).Elem()

// IsBinary 是否返回二进制类型
func IsBinary(typ reflect.Type) bool {
	if typ == nil {
		return false
	}
	if typ.Kind() == reflect.Ptr {
		return typ.Elem() == binType
	}
	return typ == binType
}

const (
	contentType        = "Content-Type"
	contentDisposition = "content-disposition"
)

// SetResponseImage 设置返回类型为图片
func SetResponseImage(writer http.ResponseWriter) {
	writer.Header().Add(contentType, "image/*")
}

// SetResponseDownload 设置返回类型为二进制
func SetResponseDownload(writer http.ResponseWriter, filename string) {
	writer.Header().Set(contentType, "application/octet-stream")
	writer.Header().Set(contentDisposition, "attachment;filename="+filename)
}
