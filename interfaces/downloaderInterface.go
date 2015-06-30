package interfaces

type DownloaderInterface interface {
	GetScheduler() SchedulerInterface
	SetScheduler(SchedulerInterface)
	Start(int)
}
