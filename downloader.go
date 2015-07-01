package singoriensis

import (
	"singoriensis/common"
	"singoriensis/interfaces"
	"time"
)

type Downloader struct {
	requests    []*Request
	scheduler   interfaces.SchedulerInterface
	middlewares []interfaces.DownloaderMiddlewareInterface
}

func NewDownloader() *Downloader {
	return &Downloader{
		middlewares: make([]interfaces.DownloaderMiddlewareInterface, 0),
	}
}

func (self *Downloader) GetScheduler() interfaces.SchedulerInterface {
	return self.scheduler
}

func (self *Downloader) SetScheduler(scheduler interfaces.SchedulerInterface) {
	self.scheduler = scheduler
}
func (self *Downloader) RegisterMiddleware(mw interfaces.DownloaderMiddlewareInterface) {
	self.middlewares = append(self.middlewares, mw)
}

func (self *Downloader) CallMiddlewareMethod(name string, params []interface{}) {
	common.CallObjMethod(self.middlewares, name, params)
}

func (self *Downloader) Start(threadNum int) {
	self.requests = make([]*Request, threadNum)

	for i := 0; i < threadNum; i++ {
		request := NewRequest()
		self.requests[i] = request
		self.requests[i].SetDelegate(self)
	}

	for i := 0; i < threadNum; i++ {
		go func(index int) {
			var urlStr string

			for {
				elem := self.scheduler.ShiftElementItem()

				if elem != nil {
					elemItem := elem.(common.ElementItem)
					urlStr = elemItem.UrlStr

					self.requests[index].Init(urlStr).Request()

					// for _, v := range self.middlewares {
					// 	v.GetData(body)
					// }

					Threads <- index

				} else {
					time.Sleep(1 * time.Second)
				}
			}

		}(i)
	}
}
