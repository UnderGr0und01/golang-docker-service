package dcontainers

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
	typeContainer "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

// NewDockerClient creates a new Docker client
func NewDockerClient() (*client.Client, error) {
	return client.NewClientWithOpts(client.FromEnv)
}

// ListContainers returns a list of all containers
func ListContainers(cli *client.Client) ([]types.Container, error) {
	return cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
}

// StartContainer starts a container by ID
func StartContainer(cli *client.Client, containerID string) error {
	return cli.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
}

// StopContainer stops a container by ID
func StopContainer(cli *client.Client, containerID string) error {
	return cli.ContainerStop(context.Background(), containerID, typeContainer.StopOptions{})
}

// GetLogs returns the logs of a container
func GetLogs(cli *client.Client, containerID string) (string, error) {
	reader, err := cli.ContainerLogs(context.Background(), containerID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
	})
	if err != nil {
		return "", err
	}
	defer reader.Close()

	logs, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}

	return string(logs), nil
}
