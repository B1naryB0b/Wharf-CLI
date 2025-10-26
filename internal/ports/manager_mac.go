//go:build darwin

package ports

func newPlatformManager() PlatformManager {
	return &macManager{}
}

type macManager struct {
	baseManager
}

func (m *macManager) GetAllPorts() (string, error) {
	return m.executeCommand("sh", "-c", "lsof -i -P -n")
}

func (m *macManager) GetPort(port string) (string, error) {
	return m.executeCommand("sh", "-c", "lsof -i :"+port+" -P -n")
}
