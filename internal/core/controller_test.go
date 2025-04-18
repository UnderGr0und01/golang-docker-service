package core

import (
	"testing"
)

func TestController(t *testing.T) {
	controller := NewController()

	// Test GetContainers
	containers, err := controller.GetContainers()
	if err != nil {
		t.Fatalf("Failed to get containers: %v", err)
	}

	// Verify containers list is not nil
	if containers == nil {
		t.Error("Containers list should not be nil")
	}

	// If there are containers, test start/stop operations
	if len(containers) > 0 {
		containerID := containers[0].ID

		// Test StopContainer
		err = controller.StopContainer(containerID)
		if err != nil {
			t.Errorf("Failed to stop container: %v", err)
		}

		// Test StartContainer
		err = controller.StartContainer(containerID)
		if err != nil {
			t.Errorf("Failed to start container: %v", err)
		}

		// Test GetLogs
		logs, err := controller.GetLogs(containerID)
		if err != nil {
			t.Errorf("Failed to get container logs: %v", err)
		}
		if logs == "" {
			t.Error("Container logs should not be empty")
		}
	}
}
