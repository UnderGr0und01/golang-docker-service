package dcontainers

import (
	"context"
	"encoding/binary"
	"io"
	"strings"

	"github.com/docker/docker/api/types"
	typeContainer "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type DContainer struct {
	ID     string `json:"ID"`
	Image  string `json:"Image"`
	State  string `json:"State"`
	Status string `json:"Status"`
}

func (container *DContainer) GetContainers() ([]DContainer, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return nil, err
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		return nil, err
	}

	containerList := make([]DContainer, len(containers))
	for i, container := range containers {
		containerList[i] = DContainer{
			ID:     container.ID,
			Status: container.Status,
			State:  container.State,
			Image:  container.Image,
		}
	}

	return containerList, nil
}

func (container *DContainer) StartContainer(id string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}

	return cli.ContainerStart(context.Background(), id, types.ContainerStartOptions{})
}

func (container *DContainer) StopContainer(id string) error {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}

	return cli.ContainerStop(context.Background(), id, typeContainer.StopOptions{})
}

func (container *DContainer) GetLogs(id string) (string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return "", err
	}

	reader, err := cli.ContainerLogs(context.Background(), id, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     false,
		Timestamps: true,
	})
	if err != nil {
		return "", err
	}
	defer reader.Close()

	header := make([]byte, 8)
	var logs strings.Builder

	for {
		_, err := reader.Read(header)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}

		size := binary.BigEndian.Uint32(header[4:8])
		if size == 0 {
			continue
		}

		message := make([]byte, size)
		_, err = reader.Read(message)
		if err != nil {
			return "", err
		}

		logs.Write(message)
		logs.WriteString("\n")
	}

	return logs.String(), nil
}
