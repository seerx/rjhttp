package rjhttp

import (
	"net/http"

	"github.com/seerx/runjson/pkg/rj"
)

// instance 默认的 builder
var instance = NewBuilder()

// Register 注册服务
func Register(loaders ...rj.Loader) *Builder {
	return instance.Register(loaders...)
}

// Inject 注册注入函数 func(arg map[string]interface{}) (*Type, error)
func Inject(fns ...interface{}) *Builder {
	for _, fn := range fns {
		if err := instance.Inject(fn); err != nil {
			panic(err)
		}
	}
	return instance
}

//InjectProxy 注册注入函数 func(ctx inject.Context) (*Type, error)
//此注入，可以直接获得 http.Request 以及 http.ResponseWriter
//func (b *Builder) InjectProxy(fns ...interface{}) error {
//	for _, fn := range fns {
//		if err := inject.Proxy(b.runner, fn); err != nil {
//			return err
//		}
//	}
//
//	return nil
//}

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

// DisableWebFileDebug 禁止调试 Web 界面文件
func DisableWebFileDebug() *Builder {
	return instance.DisableWebFileDebug()
}

// Build 创建 handler
func Build() http.Handler {
	return instance.Build()
}
