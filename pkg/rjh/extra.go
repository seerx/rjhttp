package rjh

// ExtraWriter 额外数据写入
type ExtraWriter interface {
	Write(data interface{})
}
