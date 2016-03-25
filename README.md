# singoriensis

> spider library with golang

singoriensis是参考了scrapy而编写的golang版本简单易用且高效的spider library

---

### 概述

singoriensis参考scrapy的架构，分成了downloader、scheduler、pipeliner三层，并添加了针对返回数据信息处理的process，其数据流如下：<br/>
![image](https://github.com/ErosZy/singoriensis/blob/master/image.png)


### 开始

一个简单且完整的例子如下:
```golang
    import (
        //...
        "singoriensis/common"
        //...
    )
    type MyProcess struct {}
    func (process *MyProcess) Do(page *common.Page) {
        //handle something
    }
```


```golang
    spider := sg.NewSpider("test", &MyProcess{})
    
    downloader := sg.NewDownloader()
    downloader.SetSleepTime(2 * time.Second)
    downloader.SetRetryMaxCount(0)
    downloader.RegisterMiddleware(mw.NewDefaultDownloaderMiddleware())

    scheduler := sg.NewScheduler()
    scheduler.SetUrlHeap(sg.NewDefaultUrlHeap(50))
    scheduler.RegisterMiddleware(mw.NewDefaultSchedulerMiddleware())

    pipeliner := sg.NewPipeliner()
    pipeliner.RegisterMiddleware(mw.NewDefaultPipelinerMiddleware())

    spider.SetThreadNum(1)
    spider.SetDownloader(downloader)
    spider.SetScheduler(scheduler)
    spider.SetPipeliner(pipeliner)
```