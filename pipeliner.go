package singoriensis

import (
	"github.com/ErosZy/singoriensis/common"
	"github.com/ErosZy/singoriensis/interfaces"
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
