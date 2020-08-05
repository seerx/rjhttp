package rjc

import (
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

// New 创建 http.Client
func (c *RJClient) New() (*http.Client, error) {
	var jar *cookiejar.Jar
	if c.cookies != nil && len(c.cookies) > 0 {
		uri, err := url.ParseRequestURI(c.host)
		if err != nil {
			return nil, err
		}

		cookie := []*http.Cookie{}
		for key, val := range c.cookies {
			cookie = append(cookie, &http.Cookie{
				Name:  key,
				Value: val,
			})
		}

		jar, err = cookiejar.New(nil)
		if err != nil {
			return nil, err
		}
		jar.SetCookies(uri, cookie)
	}

	tr := &http.Transport{
		//使用带超时的连接函数
		Dial: func(network, addr string) (net.Conn, error) {
			return net.DialTimeout(network, addr, c.timeout.dial)
		},
		//建立连接后读超时
		ResponseHeaderTimeout: c.timeout.response,
		DisableKeepAlives:     true,
	}

	if jar != nil {
		return &http.Client{
			Transport: tr,
			Jar:       jar,
		}, nil
	}
	return &http.Client{
		Transport: tr,
	}, nil
}
