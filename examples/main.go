package main

import (
	"fmt"
	query "github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"os"
	"regexp"
	"runtime"
	sg "singoriensis"
	cm "singoriensis/common"
	mw "singoriensis/middlewares"
	"strings"
	"sync"
	"time"
)

var fileMutex *sync.Mutex = &sync.Mutex{}
var file *os.File

type MyProcess struct {}

func (process *MyProcess) Do(page *cm.Page) {
	bytes, _ := ioutil.ReadAll(page.Res.Body)
	bodyStr := string(bytes[:])
	reader := strings.NewReader(bodyStr)
	reqUrl := page.Req.URL

	doc, _ := query.NewDocumentFromReader(reader)

	/**------------------------------波奇的产品爬虫------------------------------**/

	if reqUrl.String() == "http://shop.boqii.com/" {
		doc.Find(".menu .menu_body a").Each(func(i int, s *query.Selection) {
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
		doc.Find(".product_list .product_list_container").Each(func(i int, s *query.Selection) {
			tmp := ""
			reg := regexp.MustCompile("\\d+")
			filter := regexp.MustCompile("\\s+")

			name := filter.ReplaceAllString(s.Find(".product_name a").Text(), "")
			price := filter.ReplaceAllString(s.Find(".product_price strong").Text(), "")
			sellNum := s.Find(".product_status .product_sold").Text()
			sellNum = filter.ReplaceAllString(reg.FindString(sellNum), "")

			tmp = strings.Join([]string{tmp, sellNum, name, price}, " ")

			fileMutex.Lock()
			fmt.Println(strings.Join([]string{tmp, "\r\n"}, ""));
			file.WriteString(strings.Join([]string{tmp, "\r\n"}, ""))
			fileMutex.Unlock()
		})

		doc.Find(".pagination a").Each(func(i int, s *query.Selection) {
			href, exists := s.Attr("href")
			if exists {
				url, err := reqUrl.Parse(href)
				if err == nil {
					urlStr := url.String()
					page.AddElem(cm.NewElementItem(urlStr))
				}
			}
		})
	}

	/**------------------------------易宠的产品爬虫------------------------------**/

	//	if reqUrl.String() == "http://www.epet.com/" {
	//		doc.Find(".catelist h3 a").Each(func(i int, s *query.Selection) {
	//			href, exists := s.Attr("href")
	//			if exists {
	//				url, err := reqUrl.Parse(href)
	//				if err == nil {
	//					urlStr := url.String()
	//					page.AddElem(cm.NewElementItem(urlStr))
	//				}
	//			}
	//		})
	//	} else {
	//		count := 0
	//		doc.Find(".list_box-li").Each(func(i int, s *query.Selection) {
	//			tmp := ""
	//			reg := regexp.MustCompile("\\d+")
	//			filter := regexp.MustCompile("\\s+")
	//			sellNum := ""
	//
	//			name := filter.ReplaceAllString(s.Find(".gtitle").Text(), " ")
	//			price := filter.ReplaceAllString(s.Find(".gprice .price").Text(), " ")
	//			s.Find(".c999 em").Each(func(i int, s *query.Selection) {
	//				if i == 0 {
	//					sellNum = filter.ReplaceAllString(reg.FindString(s.Text()), " ")
	//				}
	//			})
	//
	//			tmp = strings.Join([]string{tmp, sellNum, name, price}, " ")
	//
	//			fileMutex.Lock()
	//			fmt.Println(strings.Join([]string{tmp, "\r\n"}, ""));
	//			file.WriteString(strings.Join([]string{tmp, "\r\n"}, ""))
	//			fileMutex.Unlock()
	//
	//		})
	//
	//		doc.Find(".pages a").Each(func(i int, s *query.Selection) {
	//			href, exists := s.Attr("href")
	//			if exists {
	//				url, err := reqUrl.Parse(href)
	//				if err == nil {
	//					count++
	//					urlStr := url.String()
	//					page.AddElem(cm.NewElementItem(urlStr))
	//				}
	//			}
	//		})
	//
	//		if count == 0 {
	//			fmt.Println(strings.Join([]string{reqUrl.String(), "lose pager data!!!"}, " "))
	//		}
	//	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	f, _ := os.Create("./boqi.txt")
	file = f

	spider := sg.NewSpider("test", &MyProcess{})

	downloader := sg.NewDownloader()
	downloader.SetSleepTime(2 * time.Second)
	downloader.SetRetryMaxCount(0)
	downloader.RegisterMiddleware(mw.NewDefaultDownloaderMiddleware())

	scheduler := sg.NewScheduler()
	scheduler.SetUrlHeap(sg.NewUrlHeap(50))
	scheduler.RegisterMiddleware(mw.NewDefaultSchedulerMiddleware())

	pipeliner := sg.NewPipeliner()
	pipeliner.RegisterMiddleware(mw.NewDefaultPipelinerMiddleware())

	spider.SetThreadNum(runtime.NumCPU() * 2)
	spider.SetDownloader(downloader)
	spider.SetScheduler(scheduler)
	spider.SetPipeliner(pipeliner)

	spider.AddUrl("http://shop.boqii.com/")
	//spider.AddUrl("http://www.epet.com/")

	spider.SetTimeout(5 * time.Second)
	spider.Run()
}


