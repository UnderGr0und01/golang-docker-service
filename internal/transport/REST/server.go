package REST

import (
	"docker-service/internal/core"
	"docker-service/internal/dcontainers"

	"github.com/gin-gonic/gin"
)

func StartServer() {

	// http.HandleFunc("/containers", docker.GetContainers)
	core.Handler()
	port := ":8080"
	// fmt.Printf("Server listening on %s\n", port)
	// if err := http.ListenAndServe(port, nil); err != nil {
	// 	panic(err)
	// }

	router := gin.Default()

	Docker := dcontainers.DContainer{}

	router.GET("/containers/", Docker.GetContainers)
	router.POST("/start/:id", Docker.StartContainer)
	router.POST("/stop/:id", Docker.StopContainer)

	if err := router.Run(port); err != nil {
		panic(err)
	}

}
