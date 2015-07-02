package interfaces

import (
	"singoriensis/common"
)

type ProcessInterface interface {
	Do(*common.Page)
}
