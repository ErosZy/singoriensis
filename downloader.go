package singoriensis

import (
	"fmt"
	"time"
)

type Downloader struct {
	requests  []*Request
	scheduler *Scheduler
}

func NewDownloader(scheduler *Scheduler) *Downloader {
	return &Downloader{
		scheduler: scheduler,
	}
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
				elem := self.scheduler.ShiftUrl()

				if elem != nil {
					elemItem := elem.(ElementItem)
					urlStr = elemItem.UrlStr
					method = "GET"

					bodyStr, err := self.requests[index].Init(method, urlStr).Request()

					if err == nil {
						fmt.Println(bodyStr)
					} else {
						fmt.Println(err.Error())
					}

				} else {
					time.Sleep(1 * time.Second)
				}
			}

		}(i)
	}
}
