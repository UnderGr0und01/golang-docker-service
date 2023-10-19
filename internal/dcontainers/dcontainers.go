package dcontainers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
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

// var GlobalClient = client.NewClientWithOpts(client.FromEnv) // TODO: Gloval client for docker

func (сont *DContainer) GetContainers(c *gin.Context) {
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
	JData, err := json.MarshalIndent(containerList, "", " ")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(JData))

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)

	// fmt.Fprintln(w, string(JSON))
}

func (cont *DContainer) StartContainer(c *gin.Context) {

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

func (cont *DContainer) StopContainer(c *gin.Context) {
	containerID := c.Param("id")
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = cli.ContainerStop(c.Request.Context(), containerID, container.StopOptions{}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Container %s stopped", containerID)})
}
