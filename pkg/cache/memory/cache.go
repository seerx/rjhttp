package memory

import (
	"encoding/json"
	"fmt"
	"hash/crc32"
	"time"

	"github.com/seerx/rjhttp/pkg/cache/info"
)

type item struct {
	val    interface{}
	update time.Time
	expire time.Duration
}

type Cache struct {
	defaultExpire time.Duration
	checkDuration time.Duration

	nodes     []*node
	nodeCount int
}

// NewDefault 不进行超时处理的缓存
func NewDefault() *Cache {
	return New(10, 0, 0)
}

// New 创建本机内存缓存
// nodeCount 越大性能越好，占用空间越大
// checkDuration 检查超时的时间间隔，0 或负值不进行超时检查(此时 defaultExpire 参数失效)；不宜太小，如果小于一分钟则会自动置为一分钟
// defaultExpire 默认超时时间， 0 或 负数为不限定超时
func New(nodeCount int, checkDuration time.Duration, defaultExpire time.Duration) *Cache {
	if nodeCount <= 0 {
		panic(fmt.Errorf("Node count must great then 0"))
	}
	if checkDuration > 0 && checkDuration <= time.Minute {
		checkDuration = time.Minute
	}
	nodes := make([]*node, nodeCount)
	for n := 0; n < nodeCount; n++ {
		nodes[n] = newNode(defaultExpire)
	}
	cc := &Cache{
		defaultExpire: defaultExpire,
		checkDuration: checkDuration,
		nodes:         nodes,
		nodeCount:     nodeCount,
	}
	if checkDuration > 0 {
		go cc.gc()
	}

	return cc
}

func (c *Cache) gc() {
	for _, node := range c.nodes {
		node.gc()
	}
	time.AfterFunc(c.checkDuration, func() {
		c.gc()
	})
}

func (c *Cache) Exists(key string) bool {
	index := c.hashNode(key)
	return c.nodes[index].Exists(key)
}

func (c *Cache) Get(key string) (interface{}, error) {
	index := c.hashNode(key)
	return c.nodes[index].Get(key)
}

func (c *Cache) hashNode(key string) int {
	v := int(crc32.ChecksumIEEE([]byte(key)))
	if v >= 0 {
		return v % c.nodeCount
	}
	return (-v) % c.nodeCount
}

func (c *Cache) Set(key string, value interface{}) error {
	index := c.hashNode(key)
	return c.nodes[index].Set(key, value)
}

func (c *Cache) SetX(key string, value interface{}, expire time.Duration) error {
	index := c.hashNode(key)
	return c.nodes[index].SetX(key, value, expire)
}

func (c *Cache) Remove(key string) error {
	index := c.hashNode(key)
	return c.nodes[index].Remove(key)
}

func (c *Cache) Info() string {
	cInfo := &info.CacheInfo{
		NodeCount: c.nodeCount,
		Nodes:     nil,
	}

	for _, node := range c.nodes {
		cInfo.Nodes = append(cInfo.Nodes, node.info())
	}

	data, _ := json.MarshalIndent(cInfo, "", "  ")
	return string(data)
}
