package demo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/seerx/rjhttp/pkg/rjh"

	"github.com/seerx/runjson/pkg/rj"
)

type Demo1 struct {
}

func (d *Demo1) Group() *rj.Group {
	return &grp1
}

type Req struct {
	ID      int    `json:"id" rj:"desc:ID 值,require,range:(10,100)"`
	Name    string `json:"name" rj:"desc:名称,regexp:1hdja9"`
	Bt      []bool `json:"bt" rj:"desc:测试布尔型测,deprecated"`
	Request *Req   `json:"request" rj:"desc:地轨测试"`
}

type Resp struct {
	ID   int    `json:"id" rj:"desc:ID"`
	Name string `json:"name" rj:"desc:返回名称"`
	Age  int    `json:"age" rj:"desc:年龄"`
}

func (d *Demo1) Test1(req *Req) ([]*Resp, error) {
	return []*Resp{&Resp{
		ID:   req.ID,
		Name: "Tom",
		Age:  12,
	}}, nil
}

func (d *Demo1) Test1Info() rj.FuncInfo {
	return rj.FuncInfo{
		Description: "实施函数",
		Deprecated:  false,
		History:     nil,
	}
}

func (d *Demo1) Test2(req *Req) ([]string, error) {
	return []string{"123"}, nil
}

func (d *Demo1) Test2Info() rj.FuncInfo {
	return rj.FuncInfo{
		Description:    "测试函数",
		Deprecated:     false,
		History:        nil,
		InputIsRequire: true,
	}
}

func (d Demo1) TestUploadInfo() rj.FuncInfo {
	return rj.FuncInfo{
		Description:    "测试文件上传",
		Deprecated:     false,
		InputIsRequire: false,
		History:        nil,
	}
}

func (d Demo1) TestUpload(upload *rjh.Upload) (string, error) {
	if err := upload.StoreFile("./test", "file"); err != nil {
		return "Error", err
	}
	return "ok", nil
}

func (d Demo1) TestDownload(writer http.ResponseWriter, request *http.Request) (*rjh.ResponseBinary, error) {
	for k, v := range request.Header {
		fmt.Printf("%s=%s\n", k, v[0])
	}
	rjh.SetResponseDownload(writer, "test.go")
	http.ServeFile(writer, request, "/Users/dotjava/workspace/go-projects/rjhttp/main.go")
	return nil, nil
}

func (d Demo1) TestImage(writer http.ResponseWriter) (*rjh.ResponseBinary, error) {
	//writer.Header().Add("Content-Type", "image/jpeg")
	rjh.SetResponseImage(writer)
	file, err := os.Open("/Users/dotjava/workspace/go-projects/collection/data/images/anping/andianyafan/c1/2019_07_27-2019_07_27/10.10.16.18_01_20190727191212326_TIMING.jpg")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	writer.Write(data)

	//writer
	return &rjh.ResponseBinary{}, nil
}
