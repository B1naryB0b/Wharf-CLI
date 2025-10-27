//go:build linux

package ports


func newPlatformManager() PlatformManager {
	return &linuxManager{}
}

type linuxManager struct {
	baseManager
	commonManager
}

func (m *linuxManager) OpenPort(port string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *linuxManager) OpenFirewall(port string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *linuxManager) ClosePort(port string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *linuxManager) CloseFirewall(port string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *linuxManager) StartPortLog(port string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *linuxManager) EndPortLog(port string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *linuxManager) GetActiveLogs() ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *linuxManager) GetAllPorts() (string, error) {
	return m.executeCommand("sh", "-c", "ss -tulpn")
}

func (m *linuxManager) GetPort(port string) (string, error) {
	return m.executeCommand("sh", "-c", "ss -tulpn | grep :"+port)
}
