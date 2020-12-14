package rjhttp

import (
	"net/http"

	"github.com/seerx/runjson/pkg/context"

	"github.com/seerx/runjson"

	"github.com/seerx/runjson/pkg/rj"
)

// instance 默认的 builder
var instance = NewBuilder()

// GetDefaultBuilder 获取默认 builder 实例
func GetDefaultBuilder() *Builder {
	return instance
}

// Register 注册服务
func Register(loaders ...rj.Loader) *Builder {
	return instance.Register(loaders...)
}

// Inject 注册注入函数 func(arg map[string]interface{}) (*Type, error)
func Inject(fns ...interface{}) *Builder {
	for _, fn := range fns {
		instance.Inject(fn)
	}
	return instance
}

// InjectAccessController 注册权限控制能力的注入函数
func InjectAccessController(fn interface{}) *Builder {
	return instance.InjectAccessController(fn)
}

// SetLogger 设置日志记录
func SetLogger(log context.Log) *Builder {
	return instance.SetLogger(log)
}

// LogRequest 设置打印请求数据
func LogRequest() *Builder {
	return instance.LogRequest()
}

// Before 设置运行前拦截函数
func Before(fn runjson.BeforeRun) *Builder {
	return instance.Before(fn)
}

// BeforeExecute 设置单个任务运行前拦截函数
func BeforeExecute(fn runjson.BeforeExecute) *Builder {
	return instance.BeforeExecute(fn)
}

// After 设置运行后拦截函数
func After(fn runjson.AfterRun) *Builder {
	return instance.After(fn)
}

// AfterExecute 设置单个任务运行后拦截函数
func AfterExecute(fn runjson.AfterExecute) *Builder {
	return instance.AfterExecute(fn)
}

// ErrorHandler 设置错误信息拦截函数
func ErrorHandler(fn runjson.OnError) *Builder {
	return instance.ErrorHandler(fn)
}

// EnableUpload 允许上传文件，可能影响性能，如非必要，请谨慎使用
// enable == true 时，可以使用 *runj.Upload 作为业务处理函数的参数
//			提供对传文件的操作
// 如果要上传文件操作，请在 http header 中指定业务参数字段
//	--run-json-field--
// 例如： header 中有
//		--run-json-field--: runjson
// 则 请求服务的 json 内容，应该由 POST 的 body 中的 runjson 字段提供
//		runjson=[{"service": "test.Test1", "arg": {}}]
func EnableUpload(enable bool) *Builder {
	return instance.EnableUpload(enable)
}

// SetUploadMaxSize 设置文件上传尺寸限制
func SetUploadMaxSize(maxSize int64) *Builder {
	return instance.SetUploadMaxSize(maxSize)
}

// EnableWebClient 设置是否启用 Web 界面
func EnableWebClient(enable bool) *Builder {
	return instance.EnableWebClient(enable)
}

// Build 创建 handler
func Build() http.Handler {
	return instance.Build()
}

// CreateHandler 创建 handler，与 Build 函数相同
func CreateHandler() http.Handler {
	return instance.Build()
}
