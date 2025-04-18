package dcontainers

import (
	"context"
	"testing"
	"time"

	"github.com/docker/docker/api/types"
	typeContainer "github.com/docker/docker/api/types/container"
	"github.com/stretchr/testify/assert"
)

func TestNewDockerClient(t *testing.T) {
	client, err := NewDockerClient()
	assert.NoError(t, err)
	assert.NotNil(t, client)
}

func TestListContainers(t *testing.T) {
	client, err := NewDockerClient()
	assert.NoError(t, err)

	containers, err := ListContainers(client)
	assert.NoError(t, err)
	assert.NotNil(t, containers)
}

func TestStartContainer(t *testing.T) {
	client, err := NewDockerClient()
	assert.NoError(t, err)

	// Create a unique container name using timestamp
	containerName := "test-container-" + time.Now().Format("20060102150405")

	// Create container
	container, err := client.ContainerCreate(
		context.Background(),
		&typeContainer.Config{
			Image: "alpine",
			Cmd:   []string{"echo", "test"},
		},
		nil,
		nil,
		nil,
		containerName,
	)
	assert.NoError(t, err)

	// Start container
	err = StartContainer(client, container.ID)
	assert.NoError(t, err)

	// Clean up
	err = client.ContainerRemove(context.Background(), container.ID, types.ContainerRemoveOptions{
		Force: true,
	})
	assert.NoError(t, err)
}

func TestStopContainer(t *testing.T) {
	client, err := NewDockerClient()
	assert.NoError(t, err)

	// Create a unique container name using timestamp
	containerName := "test-container-" + time.Now().Format("20060102150405")

	// Create container
	container, err := client.ContainerCreate(
		context.Background(),
		&typeContainer.Config{
			Image: "alpine",
			Cmd:   []string{"echo", "test"},
		},
		nil,
		nil,
		nil,
		containerName,
	)
	assert.NoError(t, err)

	// Start container
	err = client.ContainerStart(context.Background(), container.ID, types.ContainerStartOptions{})
	assert.NoError(t, err)

	// Stop container
	err = StopContainer(client, container.ID)
	assert.NoError(t, err)

	// Clean up
	err = client.ContainerRemove(context.Background(), container.ID, types.ContainerRemoveOptions{
		Force: true,
	})
	assert.NoError(t, err)
}

func TestGetLogs(t *testing.T) {
	client, err := NewDockerClient()
	assert.NoError(t, err)

	// Create a unique container name using timestamp
	containerName := "test-container-" + time.Now().Format("20060102150405")

	// Create container
	container, err := client.ContainerCreate(
		context.Background(),
		&typeContainer.Config{
			Image: "alpine",
			Cmd:   []string{"echo", "test"},
		},
		nil,
		nil,
		nil,
		containerName,
	)
	assert.NoError(t, err)

	// Start container
	err = client.ContainerStart(context.Background(), container.ID, types.ContainerStartOptions{})
	assert.NoError(t, err)

	// Wait for container to finish
	time.Sleep(2 * time.Second)

	// Get logs
	logs, err := GetLogs(client, container.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, logs)

	// Clean up
	err = client.ContainerRemove(context.Background(), container.ID, types.ContainerRemoveOptions{
		Force: true,
	})
	assert.NoError(t, err)
}
