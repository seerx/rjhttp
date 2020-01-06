package web

import (
	"net/http"

	"github.com/seerx/rjhttp/pkg/handlers/runj"
	"github.com/seerx/runjson"

	"github.com/seerx/rjhttp/pkg/option"
)

// WebHandler 界面处理
type WebHandler struct {
	rj         *runj.RjHandler
	hanlderMap map[string]http.Handler
}

func NewWebHandler(runner *runjson.Runner, opt *option.Options) *WebHandler {
	return &WebHandler{
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

func (w *WebHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	method := parseWebParam(request)
	if h, exists := w.hanlderMap[method]; exists {
		h.ServeHTTP(writer, request)
	} else {
		w.rj.ServeHTTP(writer, request)
	}
}
