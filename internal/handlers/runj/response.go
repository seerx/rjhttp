package runj

import "github.com/seerx/runjson/pkg/rj"

// RjResponse http 返回信息
type RjResponse struct {
	Error string      `json:"error,omitempty"`
	Extra interface{} `json:"extra,omitempty"`
	Data  rj.Response `json:"data"`
}

func (r *RjResponse) Write(data interface{}) {
	r.Extra = data
}
