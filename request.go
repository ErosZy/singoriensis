package singoriensis

import (
	"net/http"
	"net/url"
	"singoriensis/interfaces"
	"strconv"
	"strings"
	"time"
)

type Request struct {
	urlStr   string
	client   *http.Client
	params   map[string]string
	delegate interfaces.DownloaderInterface
}

type RequestError struct {
	statusCode int
}

func NewRequest(delegate interfaces.DownloaderInterface) *Request {
	client := &http.Client{
		Timeout: time.Duration(time.Millisecond * 500),
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

	self.delegate.CallMiddlewareMethod("SetRequest", []interface{}{req})

	if reqError == nil {
		res, resError := self.client.Do(req)
		if resError == nil {
			if res.StatusCode == 200 {
				self.delegate.CallMiddlewareMethod("GetResponse", []interface{}{res})
				return req, res, nil
			} else {
				err = RequestError{res.StatusCode}
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

func (err RequestError) Error() string {
	return strings.Join([]string{"status code error:", strconv.Itoa(err.statusCode)}, " ")
}
