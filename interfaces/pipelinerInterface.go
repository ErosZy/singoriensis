package interfaces

type PipelinerInterface interface {
	RegisterMiddleware(PipelinerMiddlewareInterface)
	CallMiddlewareMethod(string, []interface{})
}
