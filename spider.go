package singoriensis

import (
	"singoriensis/common"
	"singoriensis/interfaces"
	"time"
)

var Threads chan int

type Spider struct {
	threadNum  int
	taskName   string
	scheduler  interfaces.SchedulerInterface
	downloader interfaces.DownloaderInterface
}

func NewSpider(taskName string, process interface{}) *Spider {
	return &Spider{
		threadNum: 1,
		taskName:  taskName,
	}
}

func (self *Spider) SetThreadNum(num int) *Spider {
	if num <= 0 {
		panic("spider thread num can't be lt 0, please check it.")
	} else {
		self.threadNum = num
	}

	return self
}

func (self *Spider) AddUrl(urlstr string, pageType int) *Spider {

	if self.downloader == nil {
		panic("downloader instance is nil, please init downloader.")
	}

	if self.scheduler == nil {
		panic("scheduler instance is nil, please init scheduler.")
	}

	elemItem := common.ElementItem{UrlStr: urlstr, PageType: pageType}

	self.scheduler.AddElementItem(elemItem)

	return self
}

func (self *Spider) AddPipeline() *Spider {
	return self
}

func (self *Spider) SetDownloader(downloader interfaces.DownloaderInterface) *Spider {
	self.downloader = downloader
	return self
}

func (self *Spider) SetScheduler(scheduler interfaces.SchedulerInterface) *Spider {
	if self.downloader == nil {
		panic("downloader instance is nil, please init downloader.")
	}

	self.downloader.SetScheduler(scheduler)
	self.scheduler = scheduler
	return self
}

func (self *Spider) Run() {
	num := self.threadNum
	Threads = make(chan int, num)

	self.downloader.Start(num)

End:
	for {
		select {
		case <-Threads:
		case <-time.After(time.Second * 5):
			if count := self.scheduler.GetElemCount(); count == 0 {
				break End
			}
		}
	}
}
