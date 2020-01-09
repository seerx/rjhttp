package util

import (
	"fmt"
	"log"
	"net/http"

	"github.com/seerx/rjhttp"
)

// StartService 启动服务
func StartService(pattern string, port int) {
	if pattern == "" || pattern[0:1] != "/" {
		pattern = "/" + pattern
	}
	// 允许打开 API 的说明页面
	rjhttp.Default.EnableWebClient(true)
	mux := &http.ServeMux{}
	svr := &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: mux}
	mux.Handle(pattern, rjhttp.Default.Build())
	log.Fatal(svr.ListenAndServe())
}
