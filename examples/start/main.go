package main

import (
	"fmt"
	"net/http"

	"github.com/seerx/rjhttp/examples/start/demo"

	"github.com/seerx/rjhttp"
)

func init() {
	rjhttp.Default.EnableWebClient(true)
	rjhttp.Default.Register(&demo.Demo1{})
}

func main() {
	mux := &http.ServeMux{}
	svr := &http.Server{Addr: fmt.Sprintf(":%d", 8080), Handler: mux}
	mux.Handle("/rj", rjhttp.Default.Build())
	svr.ListenAndServe()
}
