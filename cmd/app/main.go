package main

import (
	"docker-service/internal/dcontainers"
	server "docker-service/internal/transport/REST"
)

func main() {

	Docker := dcontainers.DContainer{}
	server.StartServer(Docker)
}
