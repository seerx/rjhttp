package runj

import (
	"net/http"

	"github.com/seerx/runjson"

	"github.com/seerx/rjhttp/pkg/option"
)

type RjHandler struct {
	Runner *runjson.Runner
	Option *option.Configure
}

func (r *RjHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello"))
}
