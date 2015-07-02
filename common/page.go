package common

import (
	"net/http"
)

type Page struct {
	Req   *http.Request
	Res   *http.Response
	items []PipelinerItem
	elems []ElementItem
}

func NewPage(req *http.Request, res *http.Response) *Page {
	return &Page{
		Req:   req,
		Res:   res,
		items: make([]PipelinerItem, 0),
		elems: make([]ElementItem, 0),
	}
}

func (self *Page) AddItem(item PipelinerItem) {
	self.items = append(self.items, item)
}

func (self *Page) AddElem(elem ElementItem) {
	self.elems = append(self.elems, elem)
}

func (self *Page) GetAll() ([]PipelinerItem, []ElementItem) {
	return self.items, self.elems
}
