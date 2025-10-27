package ports

import "os/exec"

type baseManager struct{}

func (b *baseManager) executeCommand(shell, shellFlag, command string) (string, error) {
	cmd := exec.Command(shell, shellFlag, command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
