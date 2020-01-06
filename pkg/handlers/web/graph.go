package web

import (
	"encoding/json"
	"net/http"

	"github.com/seerx/runjson"
)

// GraphHandler API 图处理 Handler
type GraphHandler struct {
	runner *runjson.Runner
}

func (g *GraphHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	data, err := json.Marshal(g.runner.ApiInfo)
	if err != nil {
		panic(err)
	}
	writer.Header().Add("Content-Type", "application/json")
	writer.Write(data)
}
