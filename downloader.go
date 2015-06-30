package singoriensis

import (
	"fmt"
	"singoriensis/common"
	"singoriensis/interfaces"
	"time"
)

type Downloader struct {
	requests  []*Request
	scheduler interfaces.SchedulerInterface
}

func NewDownloader() *Downloader {
	return &Downloader{}
}

func (self *Downloader) GetScheduler() interfaces.SchedulerInterface {
	return self.scheduler
}

func (self *Downloader) SetScheduler(scheduler interfaces.SchedulerInterface) {
	self.scheduler = scheduler
}

func (self *Downloader) Start(threadNum int) {
	self.requests = make([]*Request, threadNum)

	for i := 0; i < threadNum; i++ {
		request := NewRequest()
		self.requests[i] = request
	}

	for i := 0; i < threadNum; i++ {
		go func(index int) {
			var urlStr string
			var method string

			for {
				elem := self.scheduler.ShiftElementItem()

				if elem != nil {
					elemItem := elem.(common.ElementItem)
					urlStr = elemItem.UrlStr
					method = "GET"

					bodyStr, err := self.requests[index].Init(method, urlStr).Request()

					if err == nil {
						fmt.Println(bodyStr)
					} else {
						fmt.Println(err.Error())
					}

					Threads <- index

				} else {
					time.Sleep(1 * time.Second)
				}
			}

		}(i)
	}
}
