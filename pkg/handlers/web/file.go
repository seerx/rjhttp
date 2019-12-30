package web

import (
	"io/ioutil"
	"net/http"
	"os"
)

const BASE = "/Users/dotjava/workspace/go-projects/rjhttp/resources/html/"

type File struct {
}

func (i *File) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fileName := request.URL.Query().Get("file")
	file, err := os.Open(BASE + fileName)
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
