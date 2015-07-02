package singoriensis

import (
	"singoriensis/common"
	"singoriensis/interfaces"
)

type Pipeliner struct {
	middlewares []interface{}
}

func NewPipeliner() *Pipeliner {
	return &Pipeliner{
		middlewares: make([]interface{}, 0),
	}
}

func (self *Pipeliner) RegisterMiddleware(mw interfaces.PipelinerMiddlewareInterface) {
	self.middlewares = append(self.middlewares, mw)
}

func (self *Pipeliner) CallMiddlewareMethod(name string, params []interface{}) {
	common.CallObjMethod(self.middlewares, name, params)
}
