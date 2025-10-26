//go:build windows

package ports

import (
	"fmt"
	"strconv"
	"strings"
)

func newPlatformManager() PlatformManager {
	return &windowsManager{}
}

type windowsManager struct {
	baseManager
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

func (m *windowsManager) PingPort(port string, timeout float64) (string, error) {
	host := "localhost"
	portNum := port

	if strings.Contains(port, ":") {
		parts := strings.Split(port, ":")
		if len(parts) == 2 {
			host = parts[0]
			portNum = parts[1]
		}
	}

	if _, err := strconv.Atoi(portNum); err != nil {
		return "", fmt.Errorf("invalid port number: %s", portNum)
	}

	timeoutSeconds := int(timeout)
	command := fmt.Sprintf("$job = Start-Job -ScriptBlock { Test-NetConnection -ComputerName %s -Port %s -InformationLevel Quiet }; Wait-Job $job -Timeout %d | Out-Null; $result = Receive-Job $job; Remove-Job $job -Force; $result", host, portNum, timeoutSeconds)
	output, err := m.executeCommand("powershell", "-Command", command)

	result := strings.TrimSpace(output)

	if result == "" {
		return fmt.Sprintf("✗ Port %s timed out on %s (timeout: %.1fs)", portNum, host, timeout), fmt.Errorf("connection timed out after %.1f seconds", timeout)
	}

	if result == "True" {
		return fmt.Sprintf("✓ Port %s is open on %s", portNum, host), nil
	}

	// Connection failed
	if err != nil {
		return fmt.Sprintf("✗ Port %s is closed on %s", portNum, host), err
	}

	return fmt.Sprintf("✗ Port %s is closed on %s", portNum, host), fmt.Errorf("port not reachable")
}

func (m *windowsManager) GetNextFreePort() (string, error) {
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
