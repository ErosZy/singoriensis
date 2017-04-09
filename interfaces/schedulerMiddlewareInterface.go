package interfaces

import (
	"github.com/ErosZy/singoriensis/common"
)

type SchedulerMiddlewareInterface interface {
	ElementItemIn(*bool, *common.ElementItem)
	ElementItemOut(*bool, *common.ElementItem)
}
