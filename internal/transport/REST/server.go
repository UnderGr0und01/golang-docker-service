package REST

import (
	"docker-service/internal/dcontainers"
	"fmt"
	"net/http"
)

func StartServer(docker dcontainers.DContainer) {

	http.HandleFunc("/containers", docker.GetContainers)

	port := ":8080"
	fmt.Printf("Server listening on %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
