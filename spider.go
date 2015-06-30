package singoriensis

const (
	HTML_TYPE = iota
	JSON_TYPE
)

type Spider struct {
	threadNum  int
	taskName   string
	scheduler  *Scheduler
	downloader *Downloader
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

	if self.downloader.scheduler == nil {
		panic("downloader's scheduler is nil, please set schduler.")
	}

	elemItem := ElementItem{UrlStr: urlstr, PageType: pageType}

	self.downloader.scheduler.AddUrl(elemItem)

	return self
}

func (self *Spider) AddPipeline() *Spider {
	return self
}

func (self *Spider) SetDownloader(downloader *Downloader) *Spider {
	self.downloader = downloader
	return self
}

func (self *Spider) Run() {
	num := self.threadNum
	self.downloader.Start(num)

	select {}
}
