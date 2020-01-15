package cache

import "time"

// Cache 缓存接口
type Cache interface {
	Exists(key string) bool
	Set(key string, value interface{}) error
	SetX(key string, value interface{}, expire time.Duration) error
	Get(key string) (interface{}, error)
	Remove(key string) error
	Info() string
}
