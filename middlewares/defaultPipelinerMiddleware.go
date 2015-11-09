package middlewares
import "fmt"

type DefaultPipelinerMiddleware struct {}

func NewDefaultPipelinerMiddleware() *DefaultPipelinerMiddleware {
	return &DefaultPipelinerMiddleware{}
}

func (self *DefaultPipelinerMiddleware) GetItems(stop *bool, items ...interface{}) {
	*stop = true
	fmt.Println(*stop, items)
}
