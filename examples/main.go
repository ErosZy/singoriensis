package main

import (
	"fmt"
	query "github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"runtime"
	sg "singoriensis"
	cm "singoriensis/common"
	mw "singoriensis/middlewares"
	"strings"
)

type MyProcess struct{}

func (process *MyProcess) Do(page *cm.Page) {
	bytes, _ := ioutil.ReadAll(page.Res.Body)
	bodyStr := string(bytes[:])
	reader := strings.NewReader(bodyStr)
	reqUrl := page.Req.URL

	fmt.Println("=============", reqUrl)

	doc, _ := query.NewDocumentFromReader(reader)

	if reqUrl.String() == "http://www.epet.com/" {
		doc.Find(".catelist h3 a").Each(func(i int, s *query.Selection) {
			href, exists := s.Attr("href")
			if exists {
				url, err := reqUrl.Parse(href)
				if err == nil {
					urlStr := url.String()
					page.AddElem(cm.NewElementItem(urlStr))
				}
			}
		})
	} else {
		doc.Find(".list_box-li .gtitle").Each(func(i int, s *query.Selection) {
			text := s.Text()
			fmt.Println(text)
		})
	}

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
		url := "http://www.epet.com/"
		spider.AddUrl(url)
	}

	spider.Run()
}
