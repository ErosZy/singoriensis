package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

type DefaultDownloaderMiddleware struct {}

func NewDefaultDownloaderMiddleware() *DefaultDownloaderMiddleware {
	return &DefaultDownloaderMiddleware{}
}

func (self *DefaultDownloaderMiddleware) SetClient(stop *bool, client *http.Client) {
	client.Timeout = 2 * time.Second
}

func (self *DefaultDownloaderMiddleware) SetRequest(stop *bool, req *http.Request) {
	req.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1")
	req.Header.Add("Accept", "text/html")
	req.Header.Add("Host", "item.m.jd.com")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Connection", "keep-alive")
}

func (self *DefaultDownloaderMiddleware) GetResponse(stop *bool, res *http.Response) {
	//fmt.Println(res)
}

func (self *DefaultDownloaderMiddleware) Error(stop *bool, client *http.Client, err error) {
	fmt.Println(err.Error())
}
