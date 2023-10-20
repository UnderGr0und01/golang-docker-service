package REST

import (
	"docker-service/internal/dcontainers"

	"docker-service/internal/core"

	"github.com/gin-gonic/gin"
)

func StartServer(Docker dcontainers.DContainer) {

	port := ":8080"

	router := gin.Default()

	core.Handler(router, Docker)

	if err := router.Run(port); err != nil {
		panic(err)
	}
}
