package ports

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

type commonManager struct{}

func (c *commonManager) PingPort(port string, timeout float64) (string, error) {
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

	address := fmt.Sprintf("%s:%s", host, portNum)
	timeoutDuration := time.Duration(timeout * float64(time.Second))

	conn, err := net.DialTimeout("tcp", address, timeoutDuration)
	if err != nil {
		var netErr net.Error
		if errors.As(err, &netErr) && netErr.Timeout() {
			return fmt.Sprintf("✗ Port %s timed out on %s (timeout: %.1fs)", portNum, host, timeout), fmt.Errorf("connection timed out after %.1f seconds", timeout)
		}
		return fmt.Sprintf("✗ Port %s is closed on %s", portNum, host), err
	}
	defer func() {
		if closeErr := conn.Close(); closeErr != nil {
			err = fmt.Errorf("failed to close connection: %w", closeErr)
		}
	}()

	return fmt.Sprintf("✓ Port %s is open on %s", portNum, host), err
}

func (c *commonManager) GetNextFreePort(count int) (string, error) {
	if count <= 0 {
		return "", fmt.Errorf("count must be greater than 0")
	}

	startPort := 1024
	maxPort := 65535

	freePorts := make([]int, 0, count)

	for port := startPort; port <= maxPort; port++ {
		if isPortFree(port) {
			freePorts = append(freePorts, port)
			if len(freePorts) == count {
				break
			}
		}
	}

	if len(freePorts) == 0 {
		return "", fmt.Errorf("no free ports found in range %d-%d", startPort, maxPort)
	}

	portStrings := make([]string, len(freePorts))
	for i, p := range freePorts {
		portStrings[i] = strconv.Itoa(p)
	}

	if len(freePorts) < count {
		return strings.Join(portStrings, "\n"), fmt.Errorf("only found %d free ports out of requested %d", len(freePorts), count)
	}

	return strings.Join(portStrings, "\n"), nil
}

func isPortFree(port int) bool {
	address := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return false
	}
	defer func() {
		_ = listener.Close()
	}()
	return true
}
