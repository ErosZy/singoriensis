package singoriensis

import (
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ErosZy/singoriensis/common"
	sErr "github.com/ErosZy/singoriensis/error"
	"github.com/ErosZy/singoriensis/interfaces"
)

type Request struct {
	urlStr   string
	client   *http.Client
	params   map[string]string
	delegate interfaces.DownloaderInterface
}

func NewRequest(delegate interfaces.DownloaderInterface) Request {
	client := &http.Client{
		Timeout: time.Duration(time.Second * 10),
	}

	delegate.CallMiddlewareMethod("SetClient", []interface{}{client})

	return Request{
		delegate: delegate,
		client:   client,
	}
}

func (self Request) Init(urlStr string) interfaces.RequestInterface {
	self.urlStr = urlStr
	return self
}

func (self Request) Request() (*common.Page, error) {
	var err interface{} = nil
	body := &strings.Reader{}

	values := url.Values{}

	if len(self.params) > 0 {
		params := self.params

		for v, k := range params {
			values.Add(v, k)
		}

		body = strings.NewReader(values.Encode())
	}

	req, reqError := http.NewRequest("GET", self.urlStr, body)

	var page *common.Page = nil

	if reqError == nil {
		requestError := sErr.RequestError{-1, false}
		self.delegate.CallMiddlewareMethod("SetRequest", []interface{}{req, &requestError})

		if requestError.Exist {
			err = requestError
		}

		res, resError := self.client.Do(req)

		if resError == nil {
			responseError := sErr.ResponseError{false}
			page := common.NewPage(req, res)
			self.delegate.CallMiddlewareMethod("GetResponse", []interface{}{page, &responseError})

			if responseError.Exist {
				err = responseError
			} else if res.StatusCode == 200 {
				return page, nil
			} else {
				err = sErr.RequestError{res.StatusCode, true}
			}
		} else {
			err = resError
		}
	} else {
		err = reqError
	}

	self.delegate.CallMiddlewareMethod("Error", []interface{}{self.client, err})

	return page, err.(error)
}
