package rjhttp

import (
	"net/http"

	"github.com/seerx/runjson/pkg/rj"

	"github.com/seerx/rjhttp/pkg/handlers/runj"

	"github.com/seerx/rjhttp/pkg/handlers/web"

	"github.com/seerx/runjson"

	"github.com/seerx/rjhttp/pkg/option"
)

type Builder struct {
	option *option.Options
	runner *runjson.Runner
}

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

// EnableWebClient 设置是否启用 Web 界面
func (b *Builder) EnableWebClient(enable bool) *Builder {
	b.option.EnableWebClient = enable
	return b
}

func (b *Builder) Build() http.Handler {
	var h http.Handler

	if b.option.EnableWebClient {
		h = web.NewWebHandler(b.runner, b.option)
	} else {
		h = runj.NewRjHandler(b.runner, b.option)
	}

	// 解析功能
	b.runner.Engage()

	return h
}
