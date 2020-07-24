package runj

import (
	"fmt"

	"github.com/seerx/runjson/pkg/rj"
)

// RjResponse http 返回信息
type RjResponse struct {
	Error string                 `json:"error,omitempty"`
	Extra map[string]interface{} `json:"extra,omitempty"`
	Data  rj.Response            `json:"data"`
}

// Set 写入数据
func (r *RjResponse) Set(key string, data interface{}) {
	if r.Extra == nil {
		r.Extra = map[string]interface{}{}
	}
	r.Extra[key] = data
}

// Get 获取数据
func (r *RjResponse) Get(key string) (interface{}, error) {
	if r.Extra != nil {
		data, ok := r.Extra[key]
		if !ok {
			return nil, fmt.Errorf("Key [%s] dosen't exists", key)
		}
		return data, nil
	}
	return nil, fmt.Errorf("Extra is empty")
}

// Remove 删除数据
func (r *RjResponse) Remove(key string) {
	if r.Extra != nil {
		delete(r.Extra, key)
	}
}

// RemoveAll 清空数据
func (r *RjResponse) RemoveAll() {
	if r.Extra != nil {
		r.Extra = nil
	}
}
