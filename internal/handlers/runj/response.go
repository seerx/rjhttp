package runj

import "github.com/seerx/runjson/pkg/rj"

// RjResponse http 返回信息
type RjResponse struct {
	Error string      `json:"error,omitempty"`
	Extra *ExtraInfo  `json:"extra,omitempty"`
	Data  rj.Response `json:"data"`
}

type ExtraInfo struct {
	Tag  int         `json:"tag"`
	Data interface{} `json:"data"`
}

func (r *RjResponse) Write(tag int, data interface{}) {
	r.Extra = &ExtraInfo{
		Tag:  tag,
		Data: data,
	}
}
