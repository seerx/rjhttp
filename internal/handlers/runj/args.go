package runj

import "net/http"

// RequestField http.Request
const RequestField = "__request__"

// WriterField http.ResponseWriter
const WriterField = "__writer__"

// MaxSizeField 最大上传文件大小
const MaxSizeField = "__upload_max_size__"

// ParseRequest 从注入函数的参数中获取 http.Request
func ParseRequest(injectArg map[string]interface{}) *http.Request {
	return injectArg[RequestField].(*http.Request)
}

// ParseWriter 从注入函数的参数中提取 http.ResponseWriter
func ParseWriter(injectArg map[string]interface{}) http.ResponseWriter {
	return injectArg[WriterField].(http.ResponseWriter)
}

// ParseUploadMaxSize 解析最大上传限制
func ParseUploadMaxSize(injectArg map[string]interface{}) int64 {
	return injectArg[MaxSizeField].(int64)
}
