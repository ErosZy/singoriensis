package middlewares
import "fmt"

type DefaultPipelinerMiddleware struct {}

func NewDefaultPipelinerMiddleware() *DefaultPipelinerMiddleware {
	return &DefaultPipelinerMiddleware{}
}

func (self *DefaultPipelinerMiddleware) GetItems(items ...interface{}) {
	fmt.Println(items)
}
