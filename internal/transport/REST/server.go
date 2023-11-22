package REST

import (
	"docker-service/internal/core"
	"sync"

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
	var wg sync.WaitGroup

	wg.Add(4)

	go func() {
		defer wg.Done()
		router.GET("/containers/", Dcontroller.GetContainers)
	}()

	go func() {
		wg.Done()
		router.POST("/start/:id", Dcontroller.StartContainer)
	}()
	go func() {
		wg.Done()
		router.POST("/stop/:id", Dcontroller.StopContainer)
	}()

	go func() {
		defer wg.Done()
		router.GET("/logs/:id", Dcontroller.GetLogs)
	}()
	//TODO:fix Race conditions
}
