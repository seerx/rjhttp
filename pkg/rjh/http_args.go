package rjh

import (
	"net/http"

	"github.com/seerx/rjhttp/internal/shuttlecraft"
)

// RequestField http.Request
const RequestField = "__request__"

// WriterField http.ResponseWriter
const WriterField = "__writer__"

// ExtraField rjh.ExtraWriter
const ExtraField = "__extra__"

// MaxSizeField 最大上传文件大小
const MaxSizeField = "__upload_max_size__"

const shuttlecraftField = "__shuttlecraft__"

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

// ParseExtra 解析额外数据管理者
func ParseExtra(injectArg map[string]interface{}) Extra {
	return injectArg[ExtraField].(Extra)
}

// ParseShuttlecraft 解析数据传递者（穿梭机）
func ParseShuttlecraft(injectArg map[string]interface{}) Shuttlecraft {
	val, ok := injectArg[shuttlecraftField]
	var sc Shuttlecraft
	if !ok {
		sc = shuttlecraft.New()
		injectArg[shuttlecraftField] = sc
	} else {
		sc, _ = val.(Shuttlecraft)
	}
	return sc
}
