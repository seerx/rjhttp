package memory

import (
	"sync"
	"time"

	"github.com/seerx/rjhttp/pkg/cache/info"
)

type node struct {
	mutex         sync.RWMutex
	data          map[string]*item
	defaultExpire time.Duration
	read          int
	write         int
	remove        int
	recycle       int
}

func newNode(defaultExpire time.Duration) *node {
	return &node{
		mutex:         sync.RWMutex{},
		data:          map[string]*item{},
		defaultExpire: defaultExpire,
	}
}

func (n *node) Exists(key string) bool {
	n.mutex.RLock()
	defer n.mutex.RUnlock()
	_, exists := n.data[key]
	return exists
}

func (n *node) Set(key string, value interface{}) error {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.data[key] = &item{
		val:    value,
		expire: n.defaultExpire,
		update: time.Now(),
	}
	n.write++
	return nil
}

func (n *node) SetX(key string, value interface{}, expire time.Duration) error {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.data[key] = &item{
		val:    value,
		expire: expire,
		update: time.Now(),
	}
	n.write++
	return nil
}

func (n *node) Get(key string) (interface{}, error) {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	val, exists := n.data[key]
	if exists {
		n.read++
		val.update = time.Now()
		return val.val, nil
	}
	return nil, nil
}

func (n *node) Remove(key string) error {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	delete(n.data, key)
	n.remove++
	return nil
}

func (n *node) gc() {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	now := time.Now()
	for key, val := range n.data {
		if val.expire <= 0 {
			continue
		}
		if val.update.Add(val.expire).Before(now) {
			delete(n.data, key)
			n.recycle++
		}
	}
}

func (n *node) info() *info.NodeInfo {
	n.mutex.RLock()
	defer n.mutex.RUnlock()
	return &info.NodeInfo{
		Count:   len(n.data),
		Read:    n.read,
		Write:   n.write,
		Remove:  n.remove,
		Recycle: n.recycle,
	}
}
