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
	pipeliner  interfaces.PipelinerInterface
	process    interfaces.ProcessInterface
}

func NewSpider(taskName string, process interfaces.ProcessInterface) *Spider {
	return &Spider{
		threadNum: 1,
		taskName:  taskName,
		process:   process,
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

func (self *Spider) AddUrl(urlstr string) *Spider {

	if self.downloader == nil {
		panic("downloader instance is nil, please init downloader.")
	}

	if self.scheduler == nil {
		panic("scheduler instance is nil, please init scheduler.")
	}

	if self.pipeliner == nil {
		panic("pipeliner instance is nil, please init pipeliner.")
	}

	elemItem := common.NewElementItem(urlstr)

	self.scheduler.AddElementItem(elemItem, false)

	return self
}

func (self *Spider) SetPipeliner(pipeliner interfaces.PipelinerInterface) *Spider {
	self.pipeliner = pipeliner
	return self
}

func (self *Spider) SetDownloader(downloader interfaces.DownloaderInterface) *Spider {
	self.downloader = downloader
	return self
}

func (self *Spider) SetScheduler(scheduler interfaces.SchedulerInterface) *Spider {
	self.scheduler = scheduler
	return self
}

func (self *Spider) Run() {
	num := self.threadNum
	Threads = make(chan int, num)

	self.downloader.SetProcess(self.process)
	self.downloader.SetPipeliner(self.pipeliner)
	self.downloader.SetScheduler(self.scheduler)
	self.downloader.Start(num)

	End:
	for {
		select {
		case <-Threads:
		case <-time.After(time.Minute * 30):
			if count := self.scheduler.GetElemCount(); count == 0 {
				break End
			}
		}
	}
}
