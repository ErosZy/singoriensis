package singoriensis

import (
	"net/http"
	"net/url"
	"singoriensis/interfaces"
	"strings"
	"time"
	sErr "singoriensis/error"
)

type Request struct {
	urlStr   string
	client   *http.Client
	params   map[string]string
	delegate interfaces.DownloaderInterface
}

func NewRequest(delegate interfaces.DownloaderInterface) *Request {
	client := &http.Client{
		Timeout: time.Duration(time.Second * 10),
	}

	delegate.CallMiddlewareMethod("SetClient", []interface{}{client})

	return &Request{
		delegate: delegate,
		client: client,
	}
}

func (self *Request) Init(urlStr string) *Request {
	self.urlStr = urlStr
	return self
}

func (self *Request) Request() (*http.Request, *http.Response, error) {
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

	if reqError == nil {
		requestError := sErr.RequestError{-1, false}
		self.delegate.CallMiddlewareMethod("SetRequest", []interface{}{req, &requestError})

		if requestError.Exist {
			err = requestError
		}

		res, resError := self.client.Do(req)

		if resError == nil {
			responseError := sErr.ResponseError{false}
			self.delegate.CallMiddlewareMethod("GetResponse", []interface{}{res, &responseError})

			if responseError.Exist {
				err = responseError
			}else if res.StatusCode == 200 {
				return req, res, nil
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

	return nil, nil, err.(error)
}
