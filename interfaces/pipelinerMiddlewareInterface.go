package interfaces

type PipelinerMiddlewareInterface interface {
	GetItems(*bool, ...interface{})
}
