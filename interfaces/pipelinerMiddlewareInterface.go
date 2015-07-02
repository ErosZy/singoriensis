package interfaces

import (
	"singoriensis/common"
)

type PipelinerMiddlewareInterface interface {
	GetItems([]common.PipelinerItem)
}
