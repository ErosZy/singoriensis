# singoriensis

> spider library with golang

singoriensis是参考了scrapy而编写的golang版本简单易用且高效的spider library

### 概述

singoriensis参考scrapy的架构，分成了downloader、scheduler、pipeliner三层，并添加了针对返回数据信息处理的process，其数据流如下：<br/>
![image](https://github.com/ErosZy/singoriensis/blob/master/img.png)


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

### 接口

singoriensis面向接口编程，这样你可以更简单的编写你自己的核心对象及其中间件，其接口定义位于interfaces文件夹下：

1. downloaderInterface : 核心下载对象
2. downloaderMiddlewareInterface : 下载中间件对象，你可以在这里处理连接代理，登陆模拟等需求
3. schedulerInterface : 核心调度对象
4. schedulerMiddlewareInterface : 调度中间件对象
5. pipelinerInterface : 核心存储流对象
6. pipelinerMiddlewareInterface : 存储中间件对象，可以按照顺序进行数据的存储，如log --> mysql --> elasticsearch
7. processInterface ：内容解析处理，解析出需要的url及存储的内容
8. urlHeapInterface : url冲重复过滤，你可以使用redis，也可以使用bloomfilter算法等对url进行重复过滤

