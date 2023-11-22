package dcontainers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/docker/docker/api/types"
	typeContainer "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

type DContainer struct {
	ID     string `json:"ID"`
	Image  string `json:"Image"`
	State  string `json:"State"`
	Status string `json:"Status"`
}

// type DContainers struct{
// 	containerList := make([]Dcontainer, len())
// }

func (container *DContainer) GetContainers(c *gin.Context) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		// panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	containerList := make([]DContainer, len(containers))

	for i, container := range containers {
		// fmt.Printf("%s %s\n", container.ID[:10], container.Image)
		containerList[i] = DContainer{
			ID:     container.ID,
			Status: container.Status,
			State:  container.State,
			Image:  container.Image,
		}
	}
	data, err := json.MarshalIndent(containerList, "", " ")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(data))

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)

	// fmt.Fprintln(w, string(JSON))
}

func (container *DContainer) StartContainer(c *gin.Context) {

	containerID := c.Param("id")
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	err = cli.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Container %s started", containerID)})
}

func (container *DContainer) StopContainer(c *gin.Context) {
	containerID := c.Param("id")
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = cli.ContainerStop(c.Request.Context(), containerID, typeContainer.StopOptions{}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Container %s stopped", containerID)})
}

func (container *DContainer) GetLogs(c *gin.Context) { //TODO:FIX Logs print
	containerID := c.Param("id")

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	reader, err := cli.ContainerLogs(c.Request.Context(), containerID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: false,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// logs, err := io.Copy(os.Stdout, reader)
	// if err != nil && err != io.EOF {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	data, err := json.MarshalIndent(reader, "", " ")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(data))

	// c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Logs %s", containerID)})
}
