package rjc

import (
	"fmt"
	"time"
)

// RJClient RunJson go 客户端
type RJClient struct {
	api     string
	host    string
	timeout struct {
		dial     time.Duration
		response time.Duration
	}
	cookies    map[string]string
	headers    map[string]string
	logRequest bool
}

// NewClient 创建客户端
// @param host 使用主机名称、域名或 IP 地址，不要带 http(s):// 前缀，也不要包含 / 符号
func NewClient(host, api string, port int) *RJClient {
	timeout := time.Duration(5 * time.Second)
	return NewClientExt(host, api, port, timeout, timeout)
	// host = fmt.Sprintf("http://%s", host)
	// if port != 80 {
	// 	api = fmt.Sprintf("%s:%d/%s", host, port, api)
	// } else {
	// 	api = fmt.Sprintf("%s/%s", host, api)
	// }
	// return &RJClient{
	// 	host: host,
	// 	api:  api,
	// 	timeout: struct {
	// 		dial     time.Duration
	// 		response time.Duration
	// 	}{
	// 		dial:     20 * time.Second,
	// 		response: 5 * time.Second,
	// 	},
	// }
}

// NewClientExt 创建客户端
func NewClientExt(host, api string, port int, dialTimeout, responseTimeout time.Duration) *RJClient {
	host = fmt.Sprintf("http://%s", host)
	if port != 80 {
		api = fmt.Sprintf("%s:%d/%s", host, port, api)
	} else {
		api = fmt.Sprintf("%s/%s", host, api)
	}
	return &RJClient{
		host: host,
		api:  api,
		timeout: struct {
			dial     time.Duration
			response time.Duration
		}{
			dial:     dialTimeout,
			response: responseTimeout,
		},
	}
}

// LogRequest 打印日志
func (c *RJClient) LogRequest() *RJClient {
	c.logRequest = true
	return c
}

// SetDialTimeout 设置连接超时，默认 20 秒
func (c *RJClient) SetDialTimeout(timeout time.Duration) *RJClient {
	c.timeout.dial = timeout
	return c
}

// SetResponseTimeout 设置反馈超时，默认 5 秒
func (c *RJClient) SetResponseTimeout(timeout time.Duration) *RJClient {
	c.timeout.response = timeout
	return c
}

// SetCookies 设置 cookie
func (c *RJClient) SetCookies(cookies map[string]string) *RJClient {
	c.cookies = cookies
	return c
}
