package rjh

// Extra 额外数据写入
type Extra interface {
	Write(key string, data interface{})
	Get(key string) (interface{}, error)
	Remove(key string)
	Clear()
}
