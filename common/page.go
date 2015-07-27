package common

import (
	"net/http"
)

type Page struct {
	Req   *http.Request
	Res   *http.Response
	items []interface{}
	elems []ElementItem
}

func NewPage(req *http.Request, res *http.Response) *Page {
	return &Page{
		Req:   req,
		Res:   res,
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
