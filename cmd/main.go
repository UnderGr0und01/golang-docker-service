package main

import (
	"docker-service/internal/dcontainers"
	server "docker-service/internal/transport/REST"
)

func main() {
	docker := new(dcontainers.DContainer)
	server.StartServer(*docker)
}
