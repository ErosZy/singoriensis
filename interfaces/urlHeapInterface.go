package interfaces

import (
	"singoriensis/common"
)

type UrlHeapInterface interface {
	Contain(common.ElementItem) bool
}
