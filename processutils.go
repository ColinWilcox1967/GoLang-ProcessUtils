package processutils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"unsafe"
)

// ProcessInfo holds the basic details of a process
type ProcessInfo struct {
	PID  int
	Name string
}

// GetAllProcesses retrieves the list of all running processes on a Windows machine.
func GetAllProcesses() ([]ProcessInfo, error) {
	// Use the tasklist command to retrieve the list of processes.
	cmd := exec.Command("tasklist")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve processes: %v", err)
	}

	lines := strings.Split(string(output), "\n")
	var processes []ProcessInfo

	// Skip the first few lines of output as they are headers
	for _, line := range lines[3:] {
		parts := strings.Fields(line)
		if len(parts) > 1 {
			processName := parts[0]
			pid := parts[1]

			pidInt, err := parsePID(pid)
			if err == nil {
				processes = append(processes, ProcessInfo{PID: pidInt, Name: processName})
			}
		}
	}
	return processes, nil
}

// IsProcessRunning checks if a process with the given name is running.
func IsProcessRunning(processName string) (bool, error) {
	processes, err := GetAllProcesses()
	if err != nil {
		return false, err
	}

	for _, process := range processes {
		if strings.EqualFold(process.Name, processName) {
			return true, nil
		}
	}
	return false, nil
}

// StopProcess stops a running process by name.
func StopProcess(processName string) error {
	cmd := exec.Command("taskkill", "/IM", processName, "/F")
	return cmd.Run()
}

// StartProcess starts a process by name.
func StartProcess(processPath string) error {
	cmd := exec.Command(processPath)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd.Start()
}

// parsePID converts a string to an integer, ignoring errors.
func parsePID(pidStr string) (int, error) {
	return strconv.Atoi(pidStr)
}
