package runj

import (
	context2 "context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/seerx/runjson/pkg/context"

	"github.com/seerx/runjson/pkg/rj"

	"github.com/seerx/runjson"

	"github.com/seerx/rjhttp/pkg/option"
)

type RjHandler struct {
	Runner  *runjson.Runner
	Option  *option.Options
	parseFn func(request *http.Request) (rj.Requests, error)
}

func NewRjHandler(runner *runjson.Runner, opt *option.Options) *RjHandler {
	if opt.EnableUpload {
		// 可以上传文件
		return &RjHandler{
			Runner:  runner,
			Option:  opt,
			parseFn: parseFieldOrBody,
		}
	}
	return &RjHandler{
		Runner:  runner,
		Option:  opt,
		parseFn: parseBody,
	}

}
func parseFieldOrBody(request *http.Request) (rj.Requests, error) {
	fieldName := request.Header.Get(option.PostFieldNameInHttpHeader)
	if fieldName == "" {
		return parseBody(request)
	}
	var reqs rj.Requests
	val := request.PostForm.Get(fieldName)
	if err := json.NewDecoder(strings.NewReader(val)).Decode(&reqs); err != nil {
		return nil, err
	}
	return reqs, nil
}

func parseBody(request *http.Request) (rj.Requests, error) {
	var reqs rj.Requests
	if err := json.NewDecoder(request.Body).Decode(&reqs); err != nil {
		return nil, err
	}
	return reqs, nil
}

func (r *RjHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	res := RjResponse{}
	reqs, err := r.parseFn(request)
	if err != nil {
		res.Error = err.Error()
	} else {
		response, err := r.Runner.RunRequests(&context.Context{
			Context: context2.Background(),
			Param: map[string]interface{}{
				requestField: request,
				writerField:  writer,
			},
		}, reqs)
		if err != nil {
			res.Error = err.Error()
		} else {
			if len(response) == 1 {
				for _, obj := range response {
					if IsBinary(obj) {
						// 不做返回操作
						return
					}
				}
			}
		}
	}

	data, err := json.Marshal(res)
	if err != nil {
		res.Error = err.Error()
	} else {
		res.Error = ""
		writer.Write(data)
	}
}
