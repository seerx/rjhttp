package rjh

// Shuttlecraft 在不同请求之间传递数据
type Shuttlecraft interface {
	// SpotCheck 读取数据，key 将在 Shuttlecraft 中保留
	Peek(string) interface{}
	// Load 装载数据到 Shuttlecraft
	Load(string, interface{})
	// Unload 从 Shuttlecraft 卸载数据， key 将从 Shuttlecraft 中删除
	Unload(string) interface{}
	// Empty 清空 Shuttlecraft
	Empty()
}
