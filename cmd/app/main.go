package main

import (
	"docker-service/internal/dcontainers"
	server "docker-service/internal/transport/REST"

	"github.com/docker/docker/client"
)

var GlobalDockerClient *client.Client

var Docker dcontainers.DContainer

func main() {

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	GlobalDockerClient = cli

	Docker = dcontainers.DContainer{}
	server.StartServer()
}
