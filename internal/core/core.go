package core

import (
	"docker-service/internal/dcontainers"
	"log"

	"github.com/gin-gonic/gin"
)

func Handler(router *gin.Engine, Docker dcontainers.DContainer) {
	log.Println("debug")
	router.GET("/containers/", Docker.GetContainers)
	router.POST("/start/:id", Docker.StartContainer)
	router.POST("/stop/:id", Docker.StopContainer)
}
