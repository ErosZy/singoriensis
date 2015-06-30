package singoriensis

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Request struct {
	method string
	urlStr string
	params map[string]string
}

type RequestError struct {
	statusCode int
}

func NewRequest() *Request {
	return &Request{}
}

func (self Request) Init(method string, urlStr string) *Request {
	self.method = method
	self.urlStr = urlStr
	return &self
}

func (self Request) Request() (string, error) {
	var err interface{} = nil
	body := &strings.Reader{}

	values := url.Values{}
	client := &http.Client{
		Timeout: time.Duration(time.Millisecond * 500),
	}

	//此处是中间件SetParams调用

	if len(self.params) > 0 {
		params := self.params

		for v, k := range params {
			values.Add(k, v)
		}

		body = strings.NewReader(values.Encode())
	}

	req, reqError := http.NewRequest(self.method, self.urlStr, body)

	//此处是中间件GetRequest调用

	if reqError == nil {
		res, resError := client.Do(req)
		if resError == nil {
			if res.StatusCode == 200 {
				//此处是中间件GetResponse调用

				bodyByte, _ := ioutil.ReadAll(res.Body)
				return string(bodyByte), nil
			} else {
				err = RequestError{res.StatusCode}
			}
		} else {
			err = resError
		}
	} else {
		err = reqError
	}

	//此处是中间件GetError调用

	return "", err.(error)
}

func (err RequestError) Error() string {
	return strings.Join([]string{"status code error:", strconv.Itoa(err.statusCode)}, " ")
}
