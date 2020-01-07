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
	option *option.Options
	runner *runjson.Runner
}

// NewBuilder 创建 builder
func NewBuilder() *Builder {
	return &Builder{
		option: &option.Options{},
		runner: runjson.New(),
	}
}

// Register 注册服务
func (b *Builder) Register(loaders ...rj.Loader) {
	b.runner.Register(loaders...)
}

// InjectRaw 注册注入函数 func(arg map[string]interface{}) (*Type, error)
func (b *Builder) Inject(fns ...interface{}) error {
	for _, fn := range fns {
		if err := b.runner.Inject(fn); err != nil {
			return err
		}
	}
	return nil
}

// Inject 注册注入函数 func(ctx inject.Context) (*Type, error)
// 此注入，可以直接获得 http.Request 以及 http.ResponseWriter
//func (b *Builder) Inject(fns ...interface{}) error {
//	for _, fn := range fns {
//		if err := inject.Proxy(b.runner, fn); err != nil {
//			return err
//		}
//	}
//
//	return nil
//}

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
