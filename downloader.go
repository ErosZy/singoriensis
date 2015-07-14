package singoriensis

import (
	"singoriensis/common"
	"singoriensis/interfaces"
	"time"
)

var retryMaxCount int

type Downloader struct {
	sleepTime   time.Duration
	requests    []*Request
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
	self.requests = make([]*Request, threadNum)

	for i := 0; i < threadNum; i++ {
		request := NewRequest()
		self.requests[i] = request
		self.requests[i].SetDelegate(self)
	}

	for i := 0; i < threadNum; i++ {
		go func(index int, retryMaxCount int) {
			var urlStr string

			for {
				elem := self.scheduler.ShiftElementItem()

				if elem != nil {
					elemItem := elem.(common.ElementItem)
					urlStr = elemItem.UrlStr

					req, res, err := self.requests[index].Init(urlStr).Request()

					if err != nil {
						if elemItem.FaildCount < retryMaxCount {
							elemItem.FaildCount += 1
							self.scheduler.AddElementItem(elemItem, true)
						}
					} else {
						params := make([]interface{}, 0)
						page := common.NewPage(req, res)
						self.process.Do(page)

						res.Body.Close()

						items, elems := page.GetAll()

						for _, v := range elems {
							self.scheduler.AddElementItem(v, false)
						}

						for _, v := range items {
							params = append(params, v)
						}

						self.pipeliner.CallMiddlewareMethod("GetItems", params)
					}

					Threads <- index

				} else {
					time.Sleep(self.sleepTime)
				}
			}

		}(i, retryMaxCount)
	}
}
