//go:build darwin

package ports

import (
	"fmt"
	"strconv"
	"strings"
)

func newPlatformManager() PlatformManager {
	return &macManager{}
}

type macManager struct {
	baseManager
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

func (m *macManager) PingPort(port string, timeout float64) (string, error) {
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

	// Use timeout command with nc (netcat) to test port connectivity
	timeoutSeconds := int(timeout)
	command := fmt.Sprintf("timeout %d nc -zv %s %s 2>&1", timeoutSeconds, host, portNum)
	output, err := m.executeCommand("sh", "-c", command)

	result := strings.TrimSpace(output)

	// nc returns 0 on success, 1 on failure, 124 on timeout
	if err != nil {
		if strings.Contains(err.Error(), "exit status 124") {
			return fmt.Sprintf("✗ Port %s timed out on %s (timeout: %.1fs)", portNum, host, timeout), fmt.Errorf("connection timed out after %.1f seconds", timeout)
		}
		return fmt.Sprintf("✗ Port %s is closed on %s", portNum, host), err
	}

	// Check if connection was successful (nc outputs "succeeded" on success)
	if strings.Contains(strings.ToLower(result), "succeeded") {
		return fmt.Sprintf("✓ Port %s is open on %s", portNum, host), nil
	}

	return fmt.Sprintf("✓ Port %s is open on %s", portNum, host), nil
}

func (m *macManager) GetNextFreePort() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m *macManager) GetAllPorts() (string, error) {
	return m.executeCommand("sh", "-c", "lsof -i -P -n")
}

func (m *macManager) GetPort(port string) (string, error) {
	return m.executeCommand("sh", "-c", "lsof -i :"+port+" -P -n")
}
