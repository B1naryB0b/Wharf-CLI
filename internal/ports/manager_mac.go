//go:build darwin

package ports

func newPlatformManager() PlatformManager {
	return &macManager{}
}

type macManager struct {
	baseManager
	commonManager
}

func (m *macManager) OpenPort(port string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *macManager) OpenFirewall(port string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *macManager) ClosePort(port string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *macManager) CloseFirewall(port string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *macManager) StartPortLog(port string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *macManager) EndPortLog(port string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *macManager) GetActiveLogs() ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *macManager) GetAllPorts() (string, error) {
	return m.executeCommand("sh", "-c", "lsof -i -P -n")
}

func (m *macManager) GetPort(port string) (string, error) {
	return m.executeCommand("sh", "-c", "lsof -i :"+port+" -P -n")
}
