package ports

type PlatformManager interface {
	GetAllPorts() (string, error)
	GetPort(port string) (string, error)

	OpenPort(port string) (string, error)
	OpenFirewall(port string) (string, error)
	ClosePort(port string) (string, error)
	CloseFirewall(port string) (string, error)

	StartPortLog(port string) (string, error)
	EndPortLog(port string) (string, error)
	GetActiveLogs() ([]string, error)
	PingPort(port string, timeout float64) (string, error)

	GetNextFreePort(count int) (string, error)
}

func NewPlatformManager() PlatformManager {
	return newPlatformManager()
}
