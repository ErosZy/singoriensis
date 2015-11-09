package interfaces

import (
	"singoriensis/common"
)

type SchedulerMiddlewareInterface interface {
	ElementItemIn(*bool, common.ElementItem)
	ElementItemOut(*bool, common.ElementItem)
}
