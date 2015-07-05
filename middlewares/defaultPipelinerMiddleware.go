package middlewares

import (
//"fmt"
	"singoriensis/common"
)

type DefaultPipelinerMiddleware struct {}

func NewDefaultPipelinerMiddleware() *DefaultPipelinerMiddleware {
	return &DefaultPipelinerMiddleware{}
}

func (self *DefaultPipelinerMiddleware) GetItems(items []common.PipelinerItem) {
	//fmt.Println(items)
}
