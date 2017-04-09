package singoriensis

import (
	"time"

	"github.com/ErosZy/singoriensis/common"
	"github.com/ErosZy/singoriensis/interfaces"
)

var retryMaxCount int

type Downloader struct {
	sleepTime   time.Duration
	requests    []interfaces.RequestInterface
	scheduler   interfaces.SchedulerInterface
	pipeliner   interfaces.PipelinerInterface
	process     interfaces.ProcessInterface
	middlewares []interfaces.DownloaderMiddlewareInterface
}

func NewDownloader() *Downloader {
	return &Downloader{
		sleepTime:   1 * time.Second,
		middlewares: make([]interfaces.DownloaderMiddlewareInterface, 0),
	}
}

func (self *Downloader) GetScheduler() interfaces.SchedulerInterface {
	return self.scheduler
}

func (self *Downloader) SetScheduler(scheduler interfaces.SchedulerInterface) {
	self.scheduler = scheduler
}

func (self *Downloader) SetPipeliner(pipeliner interfaces.PipelinerInterface) {
	self.pipeliner = pipeliner
}

func (self *Downloader) SetProcess(process interfaces.ProcessInterface) {
	self.process = process
}

func (self *Downloader) SetSleepTime(time time.Duration) {
	self.sleepTime = time
}

func (self *Downloader) SetRetryMaxCount(count int) {
	if count < 0 {
		panic("thread retry max num can't be lt 0.")
	}

	retryMaxCount = count
}

func (self *Downloader) RegisterMiddleware(mw interfaces.DownloaderMiddlewareInterface) {
	self.middlewares = append(self.middlewares, mw)
}

func (self *Downloader) CallMiddlewareMethod(name string, params []interface{}) {
	common.CallObjMethod(self.middlewares, name, params)
}

func (self *Downloader) Start(threadNum int) {
	self.requests = make([]interfaces.RequestInterface, threadNum)

	for i := 0; i < threadNum; i++ {
		request := NewRequest(self)
		self.requests[i] = request
	}

	for i := 0; i < threadNum; i++ {
		go func(index int, retryMaxCount int) {
			var urlStr string

			for {
				elem := self.scheduler.ShiftElementItem()
				if elem != nil {
					elemItem := elem.(common.ElementItem)
					urlStr = elemItem.UrlStr

					page, err := self.requests[index].Init(urlStr).Request()
					if err != nil {
						if elemItem.FaildCount < retryMaxCount {
							elemItem.FaildCount += 1
							self.scheduler.AddElementItem(elemItem, true)
						}
					} else {
						self.process.Do(page)

						items, elems := page.GetAll()

						for _, v := range elems {
							self.scheduler.AddElementItem(v, false)
						}

						if len(items) > 0 {
							self.pipeliner.CallMiddlewareMethod("GetItems", items)
						}
					}

					Threads <- index
				}

				time.Sleep(self.sleepTime)
			}

		}(i, retryMaxCount)
	}
}
