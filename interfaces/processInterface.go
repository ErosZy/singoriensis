package interfaces

import "github.com/ErosZy/singoriensis/common"

type ProcessInterface interface {
	Do(*common.Page)
}
