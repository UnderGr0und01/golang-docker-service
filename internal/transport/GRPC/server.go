package grpc

import (
	"context"
	"log"
	"net"

	"docker-service/api"
	"docker-service/internal/core"
	"docker-service/internal/models"

	"google.golang.org/grpc"
)

type Server struct {
	api.UnimplementedDockerServiceServer
	controller core.Controller
	auth       *models.Auth
}

func NewServer(controller *core.Controller, auth *models.Auth) *Server {
	return &Server{
		controller: *controller,
		auth:       auth,
	}
}

func (s *Server) Start(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	api.RegisterDockerServiceServer(grpcServer, s)

	log.Printf("Starting gRPC server on %s", port)
	return grpcServer.Serve(lis)
}

func (s *Server) Register(ctx context.Context, req *api.AuthRequest) (*api.AuthResponse, error) {
	input := &models.UserInput{
		Username: req.Username,
		Password: req.Password,
	}

	if err := s.auth.Register(input); err != nil {
		return nil, err
	}

	token, err := s.auth.GenerateToken(input.Username)
	if err != nil {
		return nil, err
	}

	return &api.AuthResponse{Token: token}, nil
}

func (s *Server) Login(ctx context.Context, req *api.AuthRequest) (*api.AuthResponse, error) {
	token, err := s.auth.Login(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	return &api.AuthResponse{Token: token}, nil
}

func (s *Server) ListContainers(ctx context.Context, _ *api.Empty) (*api.ContainerList, error) {
	containers, err := s.controller.GetContainers()
	if err != nil {
		return nil, err
	}

	protoContainers := make([]*api.Container, len(containers))
	for i, c := range containers {
		protoContainers[i] = &api.Container{
			Id:     c.ID,
			Image:  c.Image,
			State:  c.State,
			Status: c.Status,
		}
	}

	return &api.ContainerList{Containers: protoContainers}, nil
}

func (s *Server) StartContainer(ctx context.Context, req *api.ContainerID) (*api.OperationResponse, error) {
	if err := s.controller.StartContainer(req.Id); err != nil {
		return nil, err
	}

	return &api.OperationResponse{Message: "Container started successfully"}, nil
}

func (s *Server) StopContainer(ctx context.Context, req *api.ContainerID) (*api.OperationResponse, error) {
	if err := s.controller.StopContainer(req.Id); err != nil {
		return nil, err
	}

	return &api.OperationResponse{Message: "Container stopped successfully"}, nil
}

func (s *Server) GetContainerLogs(ctx context.Context, req *api.ContainerID) (*api.ContainerLogs, error) {
	logs, err := s.controller.GetLogs(req.Id)
	if err != nil {
		return nil, err
	}

	return &api.ContainerLogs{Logs: logs}, nil
}
