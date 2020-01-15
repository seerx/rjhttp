package info

// NodeInfo 节点信息
type NodeInfo struct {
	Count   int `json:"count"`
	Read    int `json:"read"`
	Write   int `json:"write"`
	Remove  int `json:"remove"`
	Recycle int `json:"recycle"`
}

// CacheInfo 缓存信息
type CacheInfo struct {
	NodeCount int         `json:"nodeCount"`
	Nodes     []*NodeInfo `json:"nodes"`
}
