package main

import (
	"fmt"

	"github.com/seerx/rjhttp"

	"github.com/seerx/rjhttp/examples/util"
	"github.com/seerx/runjson/pkg/rj"
)

type RequestArg struct {
	Name string `json:"name" rj:"desc:姓名,999,range:3<=$v<10"`
}

type Arg struct {
}

func (a *Arg) Group() *rj.Group {
	return &util.Demo
}

func (a *Arg) SayHelloInfo() rj.FuncInfo {
	return rj.FuncInfo{
		Description:    "Say hello",
		Deprecated:     false,
		InputIsRequire: true,
		History:        nil,
	}
}

func (a *Arg) SayHello(arg *RequestArg) (string, error) {
	return fmt.Sprintf("Hello %s!", arg.Name), nil
}

func main() {
	rjhttp.Register(&Arg{})
	util.StartService("", 8080)
}
