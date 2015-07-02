package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	sg "singoriensis"
	cm "singoriensis/common"
	mw "singoriensis/middlewares"
	"strings"
)

type MyProcess struct{}

func (process *MyProcess) Do(page *cm.Page) {
	bodyStr, _ := ioutil.ReadAll(page.Res.Body)
	fmt.Println(string(bodyStr))
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	spider := sg.NewSpider("test", &MyProcess{})

	downloader := sg.NewDownloader()
	downloader.SetRetryMaxCount(10)
	downloader.RegisterMiddleware(mw.NewDefaultDownloaderMiddleware())

	scheduler := sg.NewScheduler(100)
	scheduler.RegisterMiddleware(mw.NewDefaultSchedulerMiddleware())

	pipeliner := sg.NewPipeliner()
	pipeliner.RegisterMiddleware(mw.NewDefaultPipelinerMiddleware())

	spider.SetThreadNum(1)
	spider.SetDownloader(downloader)
	spider.SetScheduler(scheduler)
	spider.SetPipeliner(pipeliner)

	for i := 0; i < 1; i++ {
		url := strings.Join([]string{"http://www.baidu.com/index=", string(i)}, " ")
		spider.AddUrl(url)
	}

	spider.Run()
}
