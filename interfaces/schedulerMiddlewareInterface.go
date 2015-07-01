package interfaces

import (
	"singoriensis/common"
)

type SchedulerMiddlewareInterface interface {
	ElementItemIn(common.ElementItem)
	ElementItemOut(common.ElementItem)
}
