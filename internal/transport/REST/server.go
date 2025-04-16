package REST

import (
	"docker-service/internal/core"

	"github.com/gin-gonic/gin"
)

func StartServer(Dcontroller core.Controller) {
	port := ":8080"

	router := gin.Default()

	Handler(router, Dcontroller)

	if err := router.Run(port); err != nil {
		panic(err)
	}
}

func Handler(router *gin.Engine, Dcontroller core.Controller) {
	router.GET("/containers/", Dcontroller.GetContainers)
	router.POST("/start/:id", Dcontroller.StartContainer)
	router.POST("/stop/:id", Dcontroller.StopContainer)
	router.GET("/logs/:id", Dcontroller.GetLogs)
}
