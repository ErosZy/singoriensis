package singoriensis

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"singoriensis/interfaces"
	"strconv"
	"strings"
	"time"
)

type Request struct {
	urlStr   string
	params   map[string]string
	delegate interfaces.DownloaderInterface
}

type RequestError struct {
	statusCode int
}

func NewRequest() *Request {
	return &Request{}
}

func (self *Request) SetDelegate(delegate interfaces.DownloaderInterface) {
	self.delegate = delegate
}

func (self *Request) Init(urlStr string) *Request {
	self.urlStr = urlStr
	return self
}

func (self *Request) Request() ([]byte, error) {
	var err interface{} = nil
	body := &strings.Reader{}

	values := url.Values{}
	client := &http.Client{
		Timeout: time.Duration(time.Millisecond * 500),
	}

	self.delegate.CallMiddlewareMethod("SetClient", []interface{}{client})

	if len(self.params) > 0 {
		params := self.params

		for v, k := range params {
			values.Add(k, v)
		}

		body = strings.NewReader(values.Encode())
	}

	req, reqError := http.NewRequest("GET", self.urlStr, body)

	self.delegate.CallMiddlewareMethod("SetRequest", []interface{}{req})

	if reqError == nil {
		res, resError := client.Do(req)
		if resError == nil {
			if res.StatusCode == 200 {
				self.delegate.CallMiddlewareMethod("GetResponse", []interface{}{res})
				bodyByte, _ := ioutil.ReadAll(res.Body)
				return bodyByte, nil
			} else {
				err = RequestError{res.StatusCode}
			}
		} else {
			err = resError
		}
	} else {
		err = reqError
	}

	self.delegate.CallMiddlewareMethod("Error", []interface{}{err})

	return []byte{}, err.(error)
}

func (err RequestError) Error() string {
	return strings.Join([]string{"status code error:", strconv.Itoa(err.statusCode)}, " ")
}
