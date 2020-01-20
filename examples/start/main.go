package main

import (
	"fmt"
	"net/http"

	"github.com/seerx/rjhttp"

	"github.com/seerx/rjhttp/pkg/handlers/runj"

	"github.com/seerx/rjhttp/examples/start/demo"
)

func init() {
	rjhttp.EnableWebClient(true).
		EnableUpload(true).
		Inject(InjectResponse).
		Register(&demo.Demo1{}).
		DisableWebFileDebug()
}

func InjectResponse(arg map[string]interface{}) (http.ResponseWriter, error) {
	return runj.ParseWriter(arg), nil
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
