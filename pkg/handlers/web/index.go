package web

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/seerx/rjhttp/pkg/handlers/web/pages"
)

// Index Web 首页
type Index struct {
	debug bool
}

func returnFile(writer http.ResponseWriter, fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	writer.Write(data)
}

func (i *Index) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "text/html")

	if i.debug {
		// 调试页面
		returnFile(writer, "./resources/html/index.html")
		return
	}

	writer.Write([]byte(pages.IndexContext))
}
