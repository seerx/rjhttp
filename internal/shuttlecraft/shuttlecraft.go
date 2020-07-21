package shuttlecraft

// Shuttlecraft 数据传递实现
type Shuttlecraft struct {
	data map[string]interface{}
}

// var errNoOne = fmt.Errorf("Shuttlecraft is empty")

// New 新建
func New() *Shuttlecraft {
	return &Shuttlecraft{
		data: map[string]interface{}{},
	}
}

// Peek 检查数据，获取数据但不删除
func (s *Shuttlecraft) Peek(name string) interface{} {
	data, ok := s.data[name]
	if ok {
		return data
	}
	return nil
}

// Load 装在数据
func (s *Shuttlecraft) Load(name string, someOne interface{}) {
	s.data[name] = someOne
}

// Unload 获取数据，并且从 Shuttlecraft 删除
func (s *Shuttlecraft) Unload(name string) interface{} {
	data, ok := s.data[name]
	defer delete(s.data, name)
	if ok {
		return data
	}
	return nil
}

// Empty 清空数据
func (s *Shuttlecraft) Empty() {
	for name := range s.data {
		delete(s.data, name)
	}
}
