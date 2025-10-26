//go:build linux

package ports

func newPlatformManager() PlatformManager {
	return &linuxManager{}
}

type linuxManager struct {
	baseManager
}

func (m *linuxManager) GetAllPorts() (string, error) {
	return m.executeCommand("sh", "-c", "ss -tulpn")
}

func (m *linuxManager) GetPort(port string) (string, error) {
	return m.executeCommand("sh", "-c", "ss -tulpn | grep :"+port)
}
