package interfaces

import (
	"github.com/ErosZy/singoriensis/common"
)

type UrlHeapInterface interface {
	Contain(common.ElementItem) bool
}
