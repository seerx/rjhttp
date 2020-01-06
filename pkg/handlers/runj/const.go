package runj

import (
	"reflect"
)

const requestField = "__request__"
const writerField = "__writer__"

// TjBinary 返回二进制文件(含图片)时，返回直结构
type TjBinary struct {
	Binary string `json:"binary" rj:"desc:返回二进制数据，不要使用 ajax 请求"`
}

var binType = reflect.TypeOf((*TjBinary)(nil)).Elem()

// IsBinary 是否返回二进制类型
func IsBinary(obj interface{}) bool {
	typ := reflect.TypeOf(obj)
	if typ.Kind() == reflect.Ptr {
		return typ.Elem() == binType
	}
	return typ == binType
}
