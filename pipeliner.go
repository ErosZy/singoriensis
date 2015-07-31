package singoriensis

import (
	"singoriensis/common"
	"singoriensis/interfaces"
)

type Pipeliner struct {
	middlewares []interfaces.PipelinerMiddlewareInterface
}

func NewPipeliner() *Pipeliner {
	return &Pipeliner{
		middlewares: make([]interfaces.PipelinerMiddlewareInterface, 0),
	}
}

func (self *Pipeliner) RegisterMiddleware(mw interfaces.PipelinerMiddlewareInterface) {
	self.middlewares = append(self.middlewares, mw)
}

func (self *Pipeliner) CallMiddlewareMethod(name string, params []interface{}) {
	common.CallObjMethod(self.middlewares, name, params)
}
