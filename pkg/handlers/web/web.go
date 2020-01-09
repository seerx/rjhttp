package web

import (
	"net/http"

	"github.com/seerx/rjhttp/pkg/handlers/runj"
	"github.com/seerx/runjson"

	"github.com/seerx/rjhttp/pkg/option"
)

// Handler 界面处理
type Handler struct {
	rj         *runj.RjHandler
	hanlderMap map[string]http.Handler
}

// NewHandler 创建含有 Web 界面的 runjson handler
func NewHandler(runner *runjson.Runner, opt *option.Option) *Handler {
	return &Handler{
		rj: runj.NewRjHandler(runner, opt),
		hanlderMap: map[string]http.Handler{
			"graph": &GraphHandler{
				runner: runner,
			},
			"index": &Index{},
			"file":  &File{},
		},
	}
}

func parseWebParam(request *http.Request) string {
	if request.Method == http.MethodGet {
		return request.URL.Query().Get(option.WebParamName)
	}
	return ""
}

func (w *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	method := parseWebParam(request)
	if method == "" {
		if request.Method == http.MethodGet {
			if request.URL.RawQuery == "" {
				w.hanlderMap["index"].ServeHTTP(writer, request)
				return
			}
		}
		// 如果没有 m 参数，则认为是 runjson
		w.rj.ServeHTTP(writer, request)
	} else {
		if h, exists := w.hanlderMap[method]; exists {
			h.ServeHTTP(writer, request)
		} else {
			w.rj.ServeHTTP(writer, request)
		}
	}
}
