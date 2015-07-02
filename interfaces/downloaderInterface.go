package interfaces

type DownloaderInterface interface {
	GetScheduler() SchedulerInterface
	SetScheduler(SchedulerInterface)
	SetRetryMaxCount(int)
	SetPipeliner(PipelinerInterface)
	SetProcess(ProcessInterface)
	RegisterMiddleware(DownloaderMiddlewareInterface)
	CallMiddlewareMethod(string, []interface{})
	Start(int)
}
