package runj

import (
	context2 "context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/seerx/rjhttp/pkg/rjh"

	"github.com/seerx/runjson/pkg/context"

	"github.com/seerx/runjson/pkg/rj"

	"github.com/seerx/runjson"

	"github.com/seerx/rjhttp/pkg/option"
)

// RjHandler 处理 runjson 请求
type RjHandler struct {
	Runner  *runjson.Runner
	Option  *option.Option
	parseFn func(request *http.Request, maxSize int64) (rj.Requests, error)
}

// injectUpload 上传辅助类注入函数
func injectUpload(arg map[string]interface{}) (*rjh.Upload, error) {
	request := rjh.ParseRequest(arg)
	//writer := ParseWriter(arg)
	//maxSize := ParseUploadMaxSize(arg)
	//request.Body = http.MaxBytesReader(writer, request.Body, maxSize)
	return &rjh.Upload{Request: request}, nil
}

// NewRjHandler 创建 runjson handler
func NewRjHandler(runner *runjson.Runner, opt *option.Option) *RjHandler {
	if opt.EnableUpload {
		// 注入上传文件操作结构体
		if err := runner.Inject(injectUpload); err != nil {
			panic(err)
		}
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
func parseFieldOrBody(request *http.Request, maxSize int64) (rj.Requests, error) {
	// http POST 判断是否有指定参数 field 名称
	//	如果有则使用 PostForm 中的此值作为请求参数
	//	如果没有，则使用 body 作为请求参数

	fieldName := request.Header.Get(option.PostFieldNameInHTTPHeader)
	if fieldName == "" {
		return parseBody(request, maxSize)
	}
	var reqs rj.Requests

	if err := request.ParseMultipartForm(maxSize); err != nil {
		return nil, err
	}

	//request.ParseMultipartForm(1000000)
	//val := request.PostForm.Get(fieldName)
	val := request.FormValue(fieldName)
	if val == "" {
		return nil, fmt.Errorf("No request found")
	}
	if err := json.NewDecoder(strings.NewReader(val)).Decode(&reqs); err != nil {
		return nil, err
	}
	return reqs, nil
}

func parseBody(request *http.Request, maxSize int64) (rj.Requests, error) {
	// http POST body 作为请求参数
	var reqs rj.Requests
	if err := json.NewDecoder(request.Body).Decode(&reqs); err != nil {
		return nil, err
	}
	return reqs, nil
}

func parseQuery(request *http.Request) (rj.Requests, error) {
	// http get 请求，url 的 "?" 之后为请求内容
	// ?[{"service": "demo1.Test1", "args":{"id":1}}]
	body := request.RequestURI
	p := strings.Index(body, "?")
	if p < 0 {
		return nil, fmt.Errorf("No request param found")
	}
	query, err := url.QueryUnescape(body[p+1:])
	//query, err := url.ParseQuery(request.URL.RawQuery)
	if err != nil {
		return nil, fmt.Errorf("Request param Unescape error")
	}
	if query == "" {
		// 查询字符串是空
		return nil, fmt.Errorf("No request param found")
	}
	var reqs rj.Requests
	if err := json.NewDecoder(strings.NewReader(query)).Decode(&reqs); err != nil {
		return nil, err
	}
	return reqs, nil
}

func (r *RjHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	res := RjResponse{}
	var reqs rj.Requests
	var err error
	if request.Method == http.MethodPost || request.Method == http.MethodPut {
		// http POST PUT
		reqs, err = r.parseFn(request, r.Option.UploadMaxBytes)
	} else {
		// http Other methods
		reqs, err = parseQuery(request)
	}

	if err != nil {
		// 解析请求时发生错误
		res.Error = err.Error()
	} else {
		// 执行 runjson 请求
		response, err := r.Runner.RunRequests(&context.Context{
			Context: context2.Background(),
			Param: map[string]interface{}{
				rjh.RequestField: request,
				rjh.WriterField:  writer,
				rjh.MaxSizeField: r.Option.UploadMaxBytes,
			},
		}, reqs)
		if err != nil {
			// 发生错误
			res.Error = err.Error()
		} else {
			// 成功
			if len(response) == 1 {
				// 如果是单独的 runjson
				for _, obj := range response {
					// 如果返回的是 RjBinary 对象，说明在业务逻辑函数内部已经处理了 writer 操作
					// 此处不再需要返回任何信息，直接 return
					if rjh.IsBinary(obj) {
						// 不做返回操作
						return
					}
				}
			}
			res.Data = response
		}
	}
	// 序列化 json
	data, err := json.Marshal(res)
	if err != nil {
		// 序列化数据时发生错误
		res.Error = err.Error()
		res.Data = nil
		data, _ := json.Marshal(res)
		writer.Write(data)
	} else {
		// 返回数据
		writer.Write(data)
	}
}
