package main

import (
	"log"
	"os"

	"docker-service/internal/core"
	"docker-service/internal/models"
	grpc "docker-service/internal/transport/GRPC"
	rest "docker-service/internal/transport/REST"

	"github.com/joho/godotenv"
)

func main() {
	envFile := ".env"
	if os.Getenv("TEST_MODE") == "true" {
		envFile = ".env.test"
	}

	if err := godotenv.Load(envFile); err != nil {
		log.Printf("Warning: Error loading %s file: %v", envFile, err)
	}

	models.InitDB(false)

	sqlDB, err := models.DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection: %v", err)
	}
	defer sqlDB.Close()

	auth := models.NewAuth(models.DB)

	controller := core.NewController()

	restServer := rest.NewServer(controller, auth)
	go func() {
		if err := restServer.Start(":8081"); err != nil {
			log.Printf("REST server error: %v", err)
		}
	}()

	grpcServer := grpc.NewServer(&controller, auth)
	go func() {
		if err := grpcServer.Start(":8082"); err != nil {
			log.Printf("gRPC server error: %v", err)
		}
	}()

	select {}
}
