package middlewares

import (
//"fmt"
	"singoriensis/common"
)

type DefaultSchedulerMiddleware struct {}

func NewDefaultSchedulerMiddleware() *DefaultSchedulerMiddleware {
	return &DefaultSchedulerMiddleware{}
}

func (self *DefaultSchedulerMiddleware) ElementItemIn(stop *bool, elem *common.ElementItem) {
	//fmt.Println("in", elem)
}

func (self *DefaultSchedulerMiddleware) ElementItemOut(stop *bool, elem *common.ElementItem) {
	//fmt.Println("out", elem)
}
