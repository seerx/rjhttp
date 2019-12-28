package web

import (
	"io/ioutil"
	"net/http"
	"os"
)

type Index struct {
}

func (i *Index) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	file, err := os.Open("/Users/dotjava/workspace/vue/rjhttp/index.html")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	writer.Header().Add("Content-Type", "text/html")
	writer.Write(data)

}
