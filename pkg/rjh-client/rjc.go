package rjc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/seerx/rjhttp/pkg/option"
)

// InvokeObject 请求对象
type InvokeObject struct {
	Service string      `json:"service"`
	Args    interface{} `json:"args"`
}

// Response 反馈信息
type Response struct {
	Error string                       `json:"error"`
	Data  map[string][]*ResponseObject `json:"data"`
}

// ResponseObject 反馈项
type ResponseObject struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

// FileObject 上传的文件对象
type FileObject struct {
	Field    string
	Data     io.Reader
	FileName string
}

const argField = "body"

// Get 发送请求
func (c *RJClient) Get(fn func(res *Response) error, objs ...*InvokeObject) error {
	res, err := c.request(objs, http.MethodGet, nil)
	if err != nil {
		return err
	}
	return fn(res)
}

// GetOne 发送请求
func (c *RJClient) GetOne(obj *InvokeObject, fn func(res *ResponseObject) error) error {
	return c.Get(func(res *Response) error {
		if res.Error != "" {
			return fn(&ResponseObject{Error: res.Error})
		}
		return fn(res.Data[obj.Service][0])
	}, obj)
}

// Post 发送请求
func (c *RJClient) Post(fn func(res *Response) error, objs ...*InvokeObject) error {
	res, err := c.request(objs, http.MethodPost, nil)
	if err != nil {
		return err
	}
	return fn(res)
}

// PostOne 发送请求
func (c *RJClient) PostOne(obj *InvokeObject, fn func(res *ResponseObject) error) error {
	return c.Post(func(res *Response) error {
		if res.Error != "" {
			return fn(&ResponseObject{Error: res.Error})
		}
		return fn(res.Data[obj.Service][0])
	}, obj)
}

// Upload 上传文件
// func (c *RJClient) Upload(obj *InvokeObject, files []*FileObject, fn func(res *ResponseObject) error) error {
// 	buf := &bytes.Buffer{}
// 	writer := multipart.NewWriter(buf)
// 	// defer writer.Close()

// 	headers := map[string]string{}
// 	// 创建上传文件的内容
// 	for _, fo := range files {
// 		// part, err := writer.CreateFormFile("image", filepath.Base(dest))
// 		part, err := writer.CreateFormFile(fo.Field, fo.FileName)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = io.Copy(part, fo.Data)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	info, err := json.Marshal([]*InvokeObject{obj})
// 	if err != nil {
// 		return err
// 	}

// 	err = writer.WriteField(argField, string(info))
// 	if err != nil {
// 		return err
// 	}

// 	headers["Content-Type"] = writer.FormDataContentType()
// 	headers[option.PostFieldNameInHTTPHeader] = argField

// 	err = writer.Close()
// 	if err != nil {
// 		return err
// 	}

// 	res, err := c.doRequest(c.api, buf, http.MethodPost, headers)
// 	if err != nil {
// 		return err
// 	}
// 	if res.Error != "" {
// 		return fn(&ResponseObject{Error: res.Error})
// 	}
// 	return fn(res.Data[obj.Service][0])
// }

// Upload 上传单个文件
func (c *RJClient) Upload(obj *InvokeObject, files []*FileObject, fn func(res *ResponseObject) error) error {
	return c.UploadMutiple([]*InvokeObject{obj}, files, func(res *Response) error {
		if res.Error != "" {
			return fn(&ResponseObject{Error: res.Error})
		}
		return fn(res.Data[obj.Service][0])
	})
}

// UploadMutiple 上传文件
func (c *RJClient) UploadMutiple(requests []*InvokeObject, files []*FileObject, fn func(res *Response) error) error {
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	// defer writer.Close()

	headers := map[string]string{}
	// 创建上传文件的内容
	for _, fo := range files {
		// part, err := writer.CreateFormFile("image", filepath.Base(dest))
		part, err := writer.CreateFormFile(fo.Field, fo.FileName)
		if err != nil {
			return err
		}
		_, err = io.Copy(part, fo.Data)
		if err != nil {
			return err
		}
	}

	info, err := json.Marshal(requests)
	if err != nil {
		return err
	}

	requestInfo := string(info)
	if c.logRequest {
		fmt.Printf("send to %s:\n%s\n", c.api, requestInfo)
	}

	err = writer.WriteField(argField, requestInfo)
	if err != nil {
		return err
	}

	headers["Content-Type"] = writer.FormDataContentType()
	headers[option.PostFieldNameInHTTPHeader] = argField

	err = writer.Close()
	if err != nil {
		return err
	}

	res, err := c.doRequest(c.api, buf, http.MethodPost, headers)
	if err != nil {
		return err
	}
	if res.Error != "" {
		return fn(&Response{Error: res.Error})
	}
	return fn(res)
}

func (c *RJClient) request(data interface{}, method string, headers map[string]string) (*Response, error) {
	buf, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	// clt, err := c.New()
	// if err != nil {
	// 	return nil, err
	// }

	var requestBody io.Reader
	var addr string
	if http.MethodGet == method {
		// url.QueryEscape()
		addr = fmt.Sprintf("%s?%s", c.api, url.QueryEscape(string(buf)))
	} else {
		addr = c.api
		requestBody = bytes.NewReader(buf)
	}
	// requestInfo := string(info)
	if c.logRequest {
		fmt.Printf("send to %s:\n%s\n", addr, string(buf))
	}
	return c.doRequest(addr, requestBody, method, headers)
	// fmt.Println(addr)
	// request, err := http.NewRequest(method, addr, requestBody)
	// if err != nil {
	// 	return nil, err
	// }
	// response, err := clt.Do(request)
	// if err != nil {
	// 	return nil, err
	// }
	// defer response.Body.Close()

	// body, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	return nil, err
	// }
	// var res Response
	// if err := json.Unmarshal(body, &res); err != nil {
	// 	fmt.Println(string(body))
	// 	return nil, err
	// }
	// return &res, nil
}

// Download  下载文件
func (c *RJClient) Download(obj *InvokeObject, headers map[string]string, fn func(reader io.Reader) error) error {
	buf, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s?[%s]", c.api, url.QueryEscape(string(buf)))
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	if headers != nil {
		for key, val := range headers {
			request.Header.Add(key, val)
		}
	}

	clt, err := c.New()
	if err != nil {
		return err
	}
	response, err := clt.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return fn(response.Body)
	// _, err = io.Copy(writer, response.Body)
	// return err
}

func (c *RJClient) doRequest(url string, reader io.Reader, method string, headers map[string]string) (*Response, error) {
	// buf, err := json.Marshal(data)
	// if err != nil {
	// 	return nil, err
	// }
	clt, err := c.New()
	if err != nil {
		return nil, err
	}

	var requestBody = reader
	// var addr string
	// if http.MethodGet == method {
	// 	// url.QueryEscape()
	// 	addr = fmt.Sprintf("%s?%s", c.api, url.QueryEscape(string(buf)))
	// } else {
	// 	addr = c.api
	// 	requestBody = bytes.NewReader(buf)
	// }
	// fmt.Println(addr)
	request, err := http.NewRequest(method, url, requestBody)
	if headers != nil {
		for key, val := range headers {
			request.Header.Add(key, val)
		}
	}
	if err != nil {
		return nil, err
	}
	response, err := clt.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var res Response
	if err := json.Unmarshal(body, &res); err != nil {
		// fmt.Println(string(body))
		return nil, err
	}

	if c.logRequest {
		jd, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			fmt.Println("Error while encode response as json", err.Error())
		} else {
			fmt.Println("Response: ", string(jd))
		}
	}

	return &res, nil
}
