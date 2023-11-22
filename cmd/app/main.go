package main

import (
	"docker-service/internal/core"
	"docker-service/internal/dcontainers"
	server "docker-service/internal/transport/REST"
)

func main() {

	var Dcontroller core.Controller
	Docker := dcontainers.DContainer{}
	Dcontroller = &Docker
	server.StartServer(Dcontroller)
}
