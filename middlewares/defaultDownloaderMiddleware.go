package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

type DefaultDownloaderMiddleware struct{}

func NewDefaultDownloaderMiddleware() *DefaultDownloaderMiddleware {
	return &DefaultDownloaderMiddleware{}
}

func (self *DefaultDownloaderMiddleware) SetClient(client *http.Client) {
	client.Timeout = 1 * time.Second
}

func (self *DefaultDownloaderMiddleware) SetRequest(req *http.Request) {
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/43.0.2357.130 Safari/537.36")
	req.Header.Add("Accept", "text/html")
	req.Header.Add("Host", "www.epet.com")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en;q=0.6,ja;q=0.4")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Connection", "keep-alive")
	//fmt.Println(req)
}

func (self *DefaultDownloaderMiddleware) GetResponse(res *http.Response) {
	//fmt.Println(res)
}

func (self *DefaultDownloaderMiddleware) Error(err error) {
	fmt.Println(err.Error())
}
