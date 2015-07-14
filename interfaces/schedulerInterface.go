package interfaces

import (
	"singoriensis/common"
)

type SchedulerInterface interface {
	GetElemCount() int
	SetUrlHeap(UrlHeapInterface)
	RegisterMiddleware(SchedulerMiddlewareInterface)
	CallMiddlewareMethod(string, []interface{})
	AddElementItem(common.ElementItem, bool)
	ShiftElementItem() interface{}
}
