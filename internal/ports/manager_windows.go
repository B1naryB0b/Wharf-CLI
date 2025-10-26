//go:build windows

package ports

func newPlatformManager() PlatformManager {
	return &windowsManager{}
}

type windowsManager struct {
	baseManager
}

func (m *windowsManager) GetAllPorts() (string, error) {
	return m.executeCommand("cmd", "/C", "netstat -ano")
}

func (m *windowsManager) GetPort(port string) (string, error) {
	return m.executeCommand("cmd", "/C", "netstat -ano | findstr :"+port)
}
