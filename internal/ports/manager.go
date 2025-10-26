package ports

type PlatformManager interface {
	GetAllPorts() (string, error)
	GetPort(port string) (string, error)
}

func NewPlatformManager() PlatformManager {
	return newPlatformManager()
}
