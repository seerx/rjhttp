package web

import (
	"net/http"

	"github.com/seerx/rjhttp/pkg/handlers/web/pages"
)

// BASE HTML 文件路径
const BASE = "/Users/dotjava/workspace/go-projects/rjhttp/resources/html/"

// File 文件处理 Handler
type File struct {
}

func (i *File) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fileName := request.URL.Query().Get("file")

	//if i.debug {
	//	returnFile(writer, "./resources/html/"+fileName)
	//	return
	//}

	data, err := pages.Asset(fileName)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	//content, ok := fileMap[fileName]
	//if !ok {
	//	panic(fmt.Errorf(""))
	//}
	//file, err := os.Open(BASE + fileName)
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	//data, err := ioutil.ReadAll(file)
	//if err != nil {
	//	panic(err)
	//}
	//writer.Header().Add("Content-Type", "text/html")
	writer.Write(data)

}
