package singoriensis

import (
	"time"

	"github.com/ErosZy/singoriensis/common"
	"github.com/ErosZy/singoriensis/interfaces"
)

var Threads chan int

type Spider struct {
	threadNum  int
	taskName   string
	timeout    time.Duration
	scheduler  interfaces.SchedulerInterface
	downloader interfaces.DownloaderInterface
	pipeliner  interfaces.PipelinerInterface
	process    interfaces.ProcessInterface
}

func NewSpider(taskName string, process interfaces.ProcessInterface) *Spider {
	return &Spider{
		threadNum: 1,
		taskName:  taskName,
		timeout:   time.Second * 30,
		process:   process,
	}
}

func (self *Spider) SetThreadNum(num int) {
	if num <= 0 {
		panic("spider thread num can't be lt 0, please check it.")
	} else {
		self.threadNum = num
	}
}

func (self *Spider) AddUrl(urlstr string) {

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
}

func (self *Spider) SetPipeliner(pipeliner interfaces.PipelinerInterface) {
	self.pipeliner = pipeliner
}

func (self *Spider) SetDownloader(downloader interfaces.DownloaderInterface) {
	self.downloader = downloader
}

func (self *Spider) SetScheduler(scheduler interfaces.SchedulerInterface) {
	self.scheduler = scheduler
}

func (self *Spider) SetTimeout(time time.Duration) {
	self.timeout = time
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
		case <-time.After(self.timeout):
			if count := self.scheduler.GetElemCount(); count == 0 {
				break End
			}
		}
	}
}
