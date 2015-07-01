package middlewares

import (
	"fmt"
	"net/http"
)

type DefaultDownloaderMiddleware struct{}

func NewDefaultDownloaderMiddleware() *DefaultDownloaderMiddleware {
	return &DefaultDownloaderMiddleware{}
}

func (self *DefaultDownloaderMiddleware) SetClient(client *http.Client) {
	fmt.Println(client)
}

func (self *DefaultDownloaderMiddleware) SetRequest(req *http.Request) {
	fmt.Println(req)
}

func (self *DefaultDownloaderMiddleware) GetResponse(res *http.Response) {
	fmt.Println(res)
}

func (self *DefaultDownloaderMiddleware) GetData(data []byte) {
	fmt.Println(string(data))
}

func (self *DefaultDownloaderMiddleware) Error(err error) {
	fmt.Println(err.Error())
}
