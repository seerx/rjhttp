package runj

import "github.com/seerx/runjson/pkg/rj"

import "fmt"

// RjResponse http 返回信息
type RjResponse struct {
	Error string                 `json:"error,omitempty"`
	Extra map[string]interface{} `json:"extra,omitempty"`
	Data  rj.Response            `json:"data"`
}

// Write 写入数据
func (r *RjResponse) Write(key string, data interface{}) {
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

// Clear 清空数据
func (r *RjResponse) Clear() {
	if r.Extra != nil {
		r.Extra = nil
	}
}
