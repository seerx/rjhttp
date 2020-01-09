package web

import (
	"net/http"

	"github.com/seerx/rjhttp/pkg/handlers/web/pages"
)

// Index Web 首页
type Index struct {
}

func (i *Index) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//file, err := os.Open("/Users/dotjava/workspace/go-projects/rjhttp/resources/html/index.html")
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	//data, err := ioutil.ReadAll(file)
	//if err != nil {
	//	panic(err)
	//}
	writer.Header().Add("Content-Type", "text/html")
	writer.Write([]byte(pages.IndexContext))

}
