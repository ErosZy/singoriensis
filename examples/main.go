package main

import (
	"fmt"
	query "github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"os"
	"runtime"
	sg "singoriensis"
	cm "singoriensis/common"
	mw "singoriensis/middlewares"
	"strings"
	"sync"
)

var fileMutex *sync.Mutex = &sync.Mutex{}
var pageMutex *sync.Mutex = &sync.Mutex{}
var pageCount int = 0
var file *os.File

type MyProcess struct{}

func (process *MyProcess) Do(page *cm.Page) {
	bytes, _ := ioutil.ReadAll(page.Res.Body)
	bodyStr := string(bytes[:])
	reader := strings.NewReader(bodyStr)
	reqUrl := page.Req.URL


	pageMutex.Lock()
	pageCount++
	fmt.Println(pageCount)
	pageMutex.Unlock()


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
		count := 0
		doc.Find(".list_box-li .gtitle").Each(func(i int, s *query.Selection) {
			text := s.Text()

			fileMutex.Lock()
			file.WriteString(strings.Join([]string{text, "\n"}, ""))
			fileMutex.Unlock()

		})

		doc.Find("#gdmenu .lis_plr a").Each(func(i int, s *query.Selection) {
			href, exists := s.Attr("href")
			if exists {
				url, err := reqUrl.Parse(href)
				if err == nil {
					count++
					urlStr := url.String()
					page.AddElem(cm.NewElementItem(urlStr))
				}
			}
		})

		if count == 0 {
			fmt.Println(strings.Join([]string{reqUrl.String(), "lose pager data!!!"}, " "))
		}
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	f, _ := os.Create("./records.txt")
	file = f

	spider := sg.NewSpider("test", &MyProcess{})

	downloader := sg.NewDownloader()
	downloader.SetRetryMaxCount(10)
	downloader.RegisterMiddleware(mw.NewDefaultDownloaderMiddleware())

	scheduler := sg.NewScheduler(100)
	scheduler.RegisterMiddleware(mw.NewDefaultSchedulerMiddleware())

	pipeliner := sg.NewPipeliner()
	pipeliner.RegisterMiddleware(mw.NewDefaultPipelinerMiddleware())

	spider.SetThreadNum(runtime.NumCPU())
	spider.SetDownloader(downloader)
	spider.SetScheduler(scheduler)
	spider.SetPipeliner(pipeliner)

<<<<<<< HEAD
	spider.AddUrl("http://www.epet.com/")
=======
	for i := 0; i < 1; i++ {
		url := "http://www.epet.com/"
		spider.AddUrl(url)
	}
>>>>>>> origin/master

	spider.Run()
}
