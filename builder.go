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
	option *option.Configure
	runner *runjson.Runner
}

func NewBuilder() *Builder {
	return &Builder{
		option: &option.Configure{},
		runner: runjson.New(),
	}
}

// Register 注册服务
func (b *Builder) Register(loaders ...rj.Loader) {
	b.runner.Register(loaders...)
}

// SetWebClient 设置是否启用 Web 界面
func (b *Builder) SetWebClient(enable bool) *Builder {
	b.option.CheckWebClientTag = enable
	return b
}

func (b *Builder) Build() http.Handler {
	var h http.Handler

	if b.option.CheckWebClientTag {
		h = web.NewWebHandler(b.runner)
	} else {
		h = &runj.RjHandler{
			Runner: b.runner,
			Option: b.option,
		}
	}

	// 解析功能
	b.runner.Engage()

	return h
}
