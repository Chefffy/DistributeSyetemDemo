package registry

type Registration struct {
	ServiceName	ServiceName
	ServiceURL	string
	RequiredServices	[]ServiceName	//服务的依赖
	ServiceUpdateURL	string
	HeartbeatURL     string		//服务状态监控
}

type ServiceName string

const (
	LogService = ServiceName("LogService")
	GradingService = ServiceName("GradingService")
	PortalService  = ServiceName("Portald")
)

type patchEntry struct {
	Name ServiceName
	URL string
}

// 服务依赖变化
type patch struct {
	Added []patchEntry
	Removed []patchEntry
}