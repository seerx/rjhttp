package rjhttp

import (
	"net/http"

	"github.com/seerx/runjson/pkg/context"

	"github.com/seerx/runjson/pkg/rj"

	"github.com/seerx/rjhttp/internal/handlers/runj"

	"github.com/seerx/rjhttp/internal/handlers/web"

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

// InjectAccessController 注册权限控制相关的注入函数
func (b *Builder) InjectAccessController(fn interface{}) *Builder {
	if err := b.runner.InjectAccessController(fn); err != nil {
		panic(err)
	}
	return b
}

// Before 设置运行前拦截函数
func (b *Builder) Before(fn runjson.BeforeRun) *Builder {
	b.runner.BeforeRun(fn)
	return b
}

// BeforeExecute 设置单个任务运行前拦截函数
func (b *Builder) BeforeExecute(fn runjson.BeforeExecute) *Builder {
	b.runner.BeforeExecute(fn)
	return b
}

// After 设置运行后拦截函数
func (b *Builder) After(fn runjson.AfterRun) *Builder {
	b.runner.AfterRun(fn)
	return b
}

// AfterExecute 设置单个任务运行后拦截函数
func (b *Builder) AfterExecute(fn runjson.AfterExecute) *Builder {
	b.runner.AfterExecute(fn)
	return b
}

// ErrorHandler 设置错误信息拦截函数
func (b *Builder) ErrorHandler(fn runjson.OnError) *Builder {
	b.runner.ErrorHandler(fn)
	return b
}

// SetLogger 设置日志输出
func (b *Builder) SetLogger(log context.Log) *Builder {
	b.runner.SetLogger(log)
	return b
}

// LogRequest 设置打印请求信息
func (b *Builder) LogRequest() *Builder {
	b.option.LogRequest = true
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
