package interfaces

type DownloaderInterface interface {
	GetScheduler() SchedulerInterface
	SetScheduler(SchedulerInterface)
	RegisterMiddleware(DownloaderMiddlewareInterface)
	CallMiddlewareMethod(string, []interface{})
	Start(int)
}
