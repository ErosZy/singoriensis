package singoriensis

import (
	"container/list"
	"singoriensis/common"
	"singoriensis/interfaces"
	"sync"
)

type Scheduler struct {
	elems       *list.List
	mutex       *sync.Mutex
	middlewares []interfaces.SchedulerMiddlewareInterface
}

type SchedulerError struct{}

func NewScheduler() *Scheduler {
	return &Scheduler{
		elems:       list.New(),
		mutex:       &sync.Mutex{},
		middlewares: make([]interfaces.SchedulerMiddlewareInterface, 0),
	}
}

func (self *Scheduler) GetElemCount() int {
	return self.elems.Len()
}

func (self *Scheduler) RegisterMiddleware(mw interfaces.SchedulerMiddlewareInterface) {
	self.middlewares = append(self.middlewares, mw)
}

func (self *Scheduler) CallMiddlewareMethod(name string, params []interface{}) {
	common.CallObjMethod(self.middlewares, name, params)
}

func (self *Scheduler) AddElementItem(elem common.ElementItem) {
	self.mutex.Lock()

	self.CallMiddlewareMethod("ElementItemIn", []interface{}{elem})
	self.elems.PushBack(elem)

	self.mutex.Unlock()
}

func (self *Scheduler) ShiftElementItem() interface{} {
	var elem interface{}

	self.mutex.Lock()

	elemItem := self.elems.Front()

	if elemItem != nil {
		elem = elemItem.Value.(common.ElementItem)
		self.elems.Remove(elemItem)
	}

	self.CallMiddlewareMethod("ElementItemOut", []interface{}{elem})

	self.mutex.Unlock()

	return elem
}

func (err SchedulerError) Error() string {
	return "can't get element from scheduler."
}
