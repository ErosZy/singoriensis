package interfaces

import "github.com/ErosZy/singoriensis/common"

type SchedulerInterface interface {
	GetElemCount() int
	SetUrlHeap(UrlHeapInterface)
	RegisterMiddleware(SchedulerMiddlewareInterface)
	CallMiddlewareMethod(string, []interface{})
	AddElementItem(common.ElementItem, bool)
	ShiftElementItem() interface{}
}
