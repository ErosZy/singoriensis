package common

import (
	"net/http"
	"io/ioutil"
)

type Page struct {
	Req     *http.Request
	Res     *http.Response
	ResBody string
	items   []interface{}
	elems   []ElementItem
}

func NewPage(req *http.Request, res *http.Response) *Page {
	var bodyStr string
	bytes, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		bodyStr = ""
	}else {
		bodyStr = string(bytes)
	}

	return &Page{
		Req:   req,
		Res:   res,
		ResBody: bodyStr,
		items: make([]interface{}, 0),
		elems: make([]ElementItem, 0),
	}
}

func (self *Page) AddItem(item interface{}) {
	self.items = append(self.items, item)
}

func (self *Page) AddElem(elem ElementItem) {
	self.elems = append(self.elems, elem)
}

func (self *Page) GetAll() ([]interface{}, []ElementItem) {
	return self.items, self.elems
}
