package singoriensis

import (
	"container/list"
)

type ElementItem struct {
	UrlStr   string
	PageType int
}

type Scheduler struct {
	elems *list.List
}

type SchedulerError struct{}

func NewScheduler() *Scheduler {
	return &Scheduler{
		elems: list.New(),
	}
}

func (self *Scheduler) AddUrl(elem ElementItem) {
	self.elems.PushBack(elem)
}

func (self *Scheduler) ShiftUrl() interface{} {
	elemItem := self.elems.Front()

	if elemItem != nil {
		elem := elemItem.Value.(ElementItem)
		self.elems.Remove(elemItem)
		return elem
	}

	return nil
}

func (err SchedulerError) Error() string {
	return "can't get element from scheduler."
}
