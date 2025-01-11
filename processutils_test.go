package processutils

import (
	"testing"
)

func TestGetAllProcesses(t *testing.T) {
	processes, err := GetAllProcesses()
	if err != nil {
		t.Fatalf("Failed to get processes: %v", err)
	}
	if len(processes) == 0 {
		t.Errorf("No processes retrieved, expected at least one.")
	}
}

func TestIsProcessRunning(t *testing.T) {
	isRunning, err := IsProcessRunning("explorer.exe")
	if err != nil {
		t.Fatalf("Error checking process: %v", err)
	}
	if !isRunning {
		t.Errorf("Expected explorer.exe to be running.")
	}
}

func TestStopAndStartProcess(t *testing.T) {
	err := StartProcess("notepad.exe")
	if err != nil {
		t.Fatalf("Failed to start notepad: %v", err)
	}

	isRunning, err := IsProcessRunning("notepad.exe")
	if err != nil {
		t.Fatalf("Error checking process: %v", err)
	}
	if !isRunning {
		t.Errorf("Expected notepad.exe to be running.")
	}

	err = StopProcess("notepad.exe")
	if err != nil {
		t.Fatalf("Failed to stop notepad: %v", err)
	}

	isRunning, err = IsProcessRunning("notepad.exe")
	if err != nil {
		t.Fatalf("Error checking process: %v", err)
	}
	if isRunning {
		t.Errorf("Expected notepad.exe to be stopped.")
	}
}
