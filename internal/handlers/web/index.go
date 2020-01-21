package web

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/seerx/rjhttp/internal/handlers/web/pages"
)

// Index Web 首页
type Index struct {
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

	//if i.debug {
	//	// 调试页面
	//	returnFile(writer, "./resources/html/index.html")
	//	return
	//}

	data, err := pages.Asset("index.html")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.Write(data)
}
