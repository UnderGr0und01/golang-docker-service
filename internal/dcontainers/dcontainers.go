package dcontainers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
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

func (container *DContainer) GetContainers(w http.ResponseWriter, r *http.Request) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
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

	JSON, err := json.MarshalIndent(containerList, "", " ")

	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, string(JSON))
}
