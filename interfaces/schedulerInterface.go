package interfaces

import (
	"singoriensis/common"
)

type SchedulerInterface interface {
	GetElemCount() int
	AddElementItem(common.ElementItem)
	ShiftElementItem() interface{}
}
