package rjh

import (
	"net/http"
	"reflect"
)

// RjBinary 返回二进制文件(含图片)时，返回直结构
type RjBinary struct {
	Binary string `json:"binary" rj:"desc:返回二进制数据，由处理函数直接调用 http.ResponseWriter.Write(...) 来返回数据"`
}

var binType = reflect.TypeOf((*RjBinary)(nil)).Elem()

// IsBinary 是否返回二进制类型
func IsBinary(obj interface{}) bool {
	typ := reflect.TypeOf(obj)
	if typ.Kind() == reflect.Ptr {
		return typ.Elem() == binType
	}
	return typ == binType
}

const (
	contentType        = "Content-Type"
	contentDisposition = "content-disposition"
)

func SetResponseImage(writer http.ResponseWriter) {
	writer.Header().Add(contentType, "image/*")
}

func SetResponseDownload(writer http.ResponseWriter, filename string) {
	writer.Header().Set(contentType, "application/octet-stream")
	writer.Header().Set(contentDisposition, "attachment;filename="+filename)
}
