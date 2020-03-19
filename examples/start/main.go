package main

import (
	"fmt"
	"net/http"

	"github.com/seerx/rjhttp/pkg/rjh"
	"github.com/seerx/runjson/pkg/rj"

	"github.com/seerx/rjhttp"

	"github.com/seerx/rjhttp/examples/start/demo"
)

func init() {
	rjhttp.EnableWebClient(true).
		EnableUpload(true).
		Inject(InjectResponse).
		Inject(InjectRequest).
		Register(&demo.Demo1{})
}

// InjectResponse 注入 http.ResponseWriter 函数
func InjectResponse(arg *rj.InjectArg) (http.ResponseWriter, error) {
	return rjh.ParseWriter(arg.Args), nil
}

// InjectRequest 注入 http.Request 函数
func InjectRequest(arg *rj.InjectArg) (*http.Request, error) {
	return rjh.ParseRequest(arg.Args), nil
}

func main() {
	//fmt.Println(utils.UUID())
	//fmt.Println(utils.UUID())
	//fmt.Println(utils.UUID())
	//var cc cache.Cache
	//cc = memory.New(10, 0, 10*time.Minute)
	//
	//for n := 0; n < 100000; n++ {
	//	cc.Set(fmt.Sprintf("%d", n), n)
	//}
	//
	//for n := 0; n < 10000; n++ {
	//	cc.Get(fmt.Sprintf("%d", n))
	//}
	//
	//fmt.Println(cc.Info())

	mux := &http.ServeMux{}
	svr := &http.Server{Addr: fmt.Sprintf(":%d", 8080), Handler: mux}
	mux.Handle("/rj", rjhttp.Build())
	svr.ListenAndServe()
}
