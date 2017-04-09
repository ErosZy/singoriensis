package interfaces

import (
	"time"
)

type DownloaderInterface interface {
	GetScheduler() SchedulerInterface
	SetScheduler(SchedulerInterface)
	SetRetryMaxCount(int)
	SetPipeliner(PipelinerInterface)
	SetProcess(ProcessInterface)
	SetSleepTime(time.Duration)
	SetRequests([]RequestInterface)
	RegisterMiddleware(DownloaderMiddlewareInterface)
	CallMiddlewareMethod(string, []interface{})
	Start(int)
}
