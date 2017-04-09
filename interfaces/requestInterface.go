package interfaces

import "github.com/ErosZy/singoriensis/common"

type RequestInterface interface {
	Init(string) RequestInterface
	Request() (*common.Page, error)
}
