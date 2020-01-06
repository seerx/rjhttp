package demo

import (
	"github.com/seerx/runjson/pkg/rj"
)

type Demo1 struct {
}

func (d *Demo1) Group() *rj.Group {
	return &grp1
}

type Req struct {
	ID      int    `json:"id" rj:"desc:ID 值,require"`
	Name    string `json:"name" rj:"desc:名称"`
	Bt      []bool `json:"bt" rj:"desc:测试布尔型测,deprecated"`
	Request *Req   `json:"request" rj:"desc:地轨测试"`
}

type Resp struct {
	ID   int    `json:"id" rj:"desc:ID"`
	Name string `json:"name" rj:"desc:返回名称"`
	Age  int    `json:"age" rj:"desc:年龄"`
}

func (d *Demo1) Test1(req *Req) ([]*Resp, error) {
	return []*Resp{}, nil
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
		Description: "测试函数",
		Deprecated:  false,
		History:     nil,
	}
}
