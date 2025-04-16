package main

import (
	"docker-service/internal/config"
	"docker-service/internal/core"
	"docker-service/internal/dcontainers"
	"docker-service/internal/handlers"
	"docker-service/internal/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	log.Println("Initializing database...")
	config.InitDB()
	log.Println("Database initialized successfully")

	router := gin.Default()

	authHandler := handlers.NewAuthHandler()

	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)

	api := router.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		var Dcontroller core.Controller
		Docker := dcontainers.DContainer{}
		Dcontroller = &Docker

		api.GET("/containers/", Dcontroller.GetContainers)
		api.POST("/start/:id", Dcontroller.StartContainer)
		api.POST("/stop/:id", Dcontroller.StopContainer)
		api.GET("/logs/:id", Dcontroller.GetLogs)
	}

	log.Println("Starting server on :8081...")
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
