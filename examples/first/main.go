package main

import (
	"github.com/seerx/rjhttp"
	"github.com/seerx/rjhttp/examples/util"
	"github.com/seerx/runjson/pkg/rj"
)

type First struct {
}

func (f *First) Group() *rj.Group {
	return &util.Demo
}

// HelloInfo Hello API 的属性信息
func (f *First) HelloInfo() rj.FuncInfo {
	return rj.FuncInfo{
		Description:    "Hello world demo",
		Deprecated:     false,
		InputIsRequire: false,
		History:        nil,
	}
}

// Hello 接口 API 函数
func (f *First) Hello() (string, error) {
	return "Hello run json!", nil
}

func main() {
	rjhttp.Default.Register(&First{})
	util.StartService("", 8080)
}
