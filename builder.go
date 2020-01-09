package rjhttp

import (
	"net/http"

	"github.com/seerx/runjson/pkg/rj"

	"github.com/seerx/rjhttp/pkg/handlers/runj"

	"github.com/seerx/rjhttp/pkg/handlers/web"

	"github.com/seerx/runjson"

	"github.com/seerx/rjhttp/pkg/option"
)

// Builder 定义
type Builder struct {
	option *option.Option
	runner *runjson.Runner
}

// NewBuilder 创建 builder
func NewBuilder() *Builder {
	return &Builder{
		option: option.NewOption(),
		runner: runjson.New(),
	}
}

// Register 注册服务
func (b *Builder) Register(loaders ...rj.Loader) *Builder {
	b.runner.Register(loaders...)
	return b
}

// Inject 注册注入函数 func(arg map[string]interface{}) (*Type, error)
func (b *Builder) Inject(fns ...interface{}) *Builder {
	for _, fn := range fns {
		if err := b.runner.Inject(fn); err != nil {
			panic(err)
		}
	}
	return b
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
func (b *Builder) EnableUpload(enable bool) *Builder {
	b.option.EnableUpload = enable
	return b
}

// SetUploadMaxSize 设置文件上传尺寸限制
func (b *Builder) SetUploadMaxSize(maxSize int64) *Builder {
	b.option.UploadMaxBytes = maxSize
	return b
}

// EnableWebClient 设置是否启用 Web 界面
func (b *Builder) EnableWebClient(enable bool) *Builder {
	b.option.EnableWebClient = enable
	return b
}

// DisableWebFileDebug 禁止调试 Web 界面文件
func (b *Builder) DisableWebFileDebug() *Builder {
	b.option.WebDebug = false
	return b
}

// Build 创建 handler
func (b *Builder) Build() http.Handler {
	var h http.Handler

	if b.option.EnableWebClient {
		h = web.NewHandler(b.runner, b.option)
	} else {
		h = runj.NewRjHandler(b.runner, b.option)
	}

	// 解析功能
	b.runner.Engage()

	return h
}
