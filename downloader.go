package singoriensis

import (
	"reflect"
	"singoriensis/common"
	"singoriensis/interfaces"
	"time"
)

type Downloader struct {
	requests   []*Request
	scheduler  interfaces.SchedulerInterface
	middleware []interfaces.DownloaderMiddlewareInterface
}

func NewDownloader() *Downloader {
	return &Downloader{
		middleware: make([]interfaces.DownloaderMiddlewareInterface, 0),
	}
}

func (self *Downloader) GetScheduler() interfaces.SchedulerInterface {
	return self.scheduler
}

func (self *Downloader) SetScheduler(scheduler interfaces.SchedulerInterface) {
	self.scheduler = scheduler
}
func (self *Downloader) RegisterMiddleware(mw interfaces.DownloaderMiddlewareInterface) {
	self.middleware = append(self.middleware, mw)
}

func (self *Downloader) CallMiddlewareMethod(name string, params []interface{}) {
	in := make([]reflect.Value, 0)

	for _, v := range params {
		in = append(in, reflect.ValueOf(v))
	}

	for _, v := range self.middleware {
		value := reflect.ValueOf(v)
		method := value.MethodByName(name)
		if method.IsValid() {
			method.Call(in)
		}
	}
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
			var method string

			for {
				elem := self.scheduler.ShiftElementItem()

				if elem != nil {
					elemItem := elem.(common.ElementItem)
					urlStr = elemItem.UrlStr
					method = "GET"

					body, _ := self.requests[index].Init(method, urlStr).Request()

					for _, v := range self.middleware {
						v.GetData(body)
					}

					Threads <- index

				} else {
					time.Sleep(1 * time.Second)
				}
			}

		}(i)
	}
}
