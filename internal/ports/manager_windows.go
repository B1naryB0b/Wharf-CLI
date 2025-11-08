//go:build windows

package ports

func newPlatformManager() PlatformManager {
	return &windowsManager{}
}

type windowsManager struct {
	baseManager
	commonManager
}

func (m *windowsManager) OpenPort(port string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *windowsManager) OpenFirewall(port string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *windowsManager) ClosePort(port string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *windowsManager) CloseFirewall(port string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *windowsManager) StartPortLog(port string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *windowsManager) EndPortLog(port string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *windowsManager) GetActiveLogs() ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *windowsManager) WatchPort(port string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *windowsManager) GetAllPorts() (string, error) {
	return m.executeCommand("cmd", "/C", "netstat -ano")
}

func (m *windowsManager) GetPort(port string) (string, error) {
	return m.executeCommand("cmd", "/C", "netstat -ano | findstr :"+port)
}
