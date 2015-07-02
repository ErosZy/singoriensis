package main

import (
	sg "singoriensis"
	mw "singoriensis/middlewares"
)

func main() {
	spider := sg.NewSpider("test", nil)

	downloader := sg.NewDownloader()
	downloader.RegisterMiddleware(mw.NewDefaultDownloaderMiddleware())

	scheduler := sg.NewScheduler(100)
	scheduler.RegisterMiddleware(mw.NewDefaultSchedulerMiddleware())

	spider.SetThreadNum(1).SetDownloader(downloader).SetScheduler(scheduler)

	for i := 0; i < 1000; i++ {
		spider.AddUrl("http://www.baidu.com/")
	}

	spider.Run()
}
