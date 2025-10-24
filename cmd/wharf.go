package cmd

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
)

func ListPorts(port string) (result string) {
	if port != "" {
		return GetSpecifiedPort(port)
	} else {
		return GetAllPorts()
	}
}

func GetSpecifiedPort(port string) (result string) {
	if _, err := strconv.Atoi(port); err != nil {
		fmt.Printf("Error: Invalid port number '%s'\n", port)
		return
	}

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", fmt.Sprintf("netstat -ano | findstr :%s", port))
	case "linux":
		cmd = exec.Command("sh", "-c", fmt.Sprintf("ss -tulpn | grep :%s", port))
	case "darwin":
		cmd = exec.Command("lsof", "-i", fmt.Sprintf(":%s", port), "-P", "-n")
	default:
		fmt.Printf("Error: Platform %s is not supported\n", runtime.GOOS)
		return
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("No process found on port %s\n", port)
		return
	}

	return string(output)
}

func GetAllPorts() (result string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("netstat", "-ano")
	case "linux":
		cmd = exec.Command("ss", "-tulpn")
	case "darwin":
		cmd = exec.Command("lsof", "-i", "-P", "-n")
	default:
		fmt.Printf("Error: Platform %s is not supported\n", runtime.GOOS)
		return
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: Failed to execute command - %v\n", err)
		return
	}

	return string(output)
}
