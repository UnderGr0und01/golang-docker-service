package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"docker-service/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr = flag.String("addr", "localhost:8082", "The server address in the format of host:port")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := api.NewDockerServiceClient(conn)

	registerResp, err := client.Register(context.Background(), &api.AuthRequest{
		Username: "grpcuser",
		Password: "grpcpass123",
	})
	if err != nil {
		log.Printf("Registration failed: %v", err)
	} else {
		log.Printf("Registration successful, token: %s", registerResp.Token)
	}

	loginResp, err := client.Login(context.Background(), &api.AuthRequest{
		Username: "grpcuser",
		Password: "grpcpass123",
	})
	if err != nil {
		log.Printf("Login failed: %v", err)
		os.Exit(1)
	}
	log.Printf("Login successful, token: %s", loginResp.Token)

	ctx := context.WithValue(context.Background(), "token", loginResp.Token)

	containers, err := client.ListContainers(ctx, &api.Empty{})
	if err != nil {
		log.Printf("Failed to list containers: %v", err)
		os.Exit(1)
	}
	fmt.Println("\nContainers:")
	for _, c := range containers.Containers {
		fmt.Printf("ID: %s, Image: %s, State: %s, Status: %s\n", c.Id, c.Image, c.State, c.Status)
	}

	if len(containers.Containers) > 0 {
		logs, err := client.GetContainerLogs(ctx, &api.ContainerID{Id: containers.Containers[0].Id})
		if err != nil {
			log.Printf("Failed to get container logs: %v", err)
		} else {
			fmt.Println("\nContainer logs:")
			fmt.Println(logs.Logs)
		}
	}

	if len(containers.Containers) > 0 {
		stopResp, err := client.StopContainer(ctx, &api.ContainerID{Id: containers.Containers[0].Id})
		if err != nil {
			log.Printf("Failed to stop container: %v", err)
		} else {
			fmt.Printf("\nStop container response: %s\n", stopResp.Message)
			time.Sleep(2 * time.Second)
		}

		startResp, err := client.StartContainer(ctx, &api.ContainerID{Id: containers.Containers[0].Id})
		if err != nil {
			log.Printf("Failed to start container: %v", err)
		} else {
			fmt.Printf("Start container response: %s\n", startResp.Message)
		}
	}
}
