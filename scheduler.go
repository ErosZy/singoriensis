package singoriensis

import (
	"container/list"
	"singoriensis/common"
	"sync"
)

type Scheduler struct {
	elems *list.List
	mutex *sync.Mutex
}

type SchedulerError struct{}

func NewScheduler() *Scheduler {
	return &Scheduler{
		elems: list.New(),
		mutex: &sync.Mutex{},
	}
}

func (self *Scheduler) GetElemCount() int {
	return self.elems.Len()
}

func (self *Scheduler) AddElementItem(elem common.ElementItem) {
	self.mutex.Lock()
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

	self.mutex.Unlock()

	return elem
}

func (err SchedulerError) Error() string {
	return "can't get element from scheduler."
}
