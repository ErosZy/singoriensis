package interfaces

import (
	"singoriensis/common"
)

type SchedulerInterface interface {
	GetElemCount() int
	RegisterMiddleware(SchedulerMiddlewareInterface)
	CallMiddlewareMethod(string, []interface{})
	AddElementItem(common.ElementItem)
	ShiftElementItem() interface{}
}
