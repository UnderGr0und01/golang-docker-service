package rest

import (
	"log"
	"net/http"

	"docker-service/internal/core"
	"docker-service/internal/middleware"
	"docker-service/internal/models"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router     *gin.Engine
	controller core.Controller
	auth       *models.Auth
}

func NewServer(controller core.Controller, auth *models.Auth) *Server {
	router := gin.Default()
	server := &Server{
		router:     router,
		controller: controller,
		auth:       auth,
	}
	server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() {
	auth := s.router.Group("/api/auth")
	{
		auth.POST("/register", s.handleRegister)
		auth.POST("/login", s.handleLogin)
	}

	api := s.router.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/containers", s.handleGetContainers)
		api.POST("/containers/:id/start", s.handleStartContainer)
		api.POST("/containers/:id/stop", s.handleStopContainer)
		api.GET("/containers/:id/logs", s.handleGetLogs)
	}
}

func (s *Server) handleGetContainers(c *gin.Context) {
	containers, err := s.controller.GetContainers()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, containers)
}

func (s *Server) handleStartContainer(c *gin.Context) {
	id := c.Param("id")
	if err := s.controller.StartContainer(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Container started successfully"})
}

func (s *Server) handleStopContainer(c *gin.Context) {
	id := c.Param("id")
	if err := s.controller.StopContainer(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Container stopped successfully"})
}

func (s *Server) handleGetLogs(c *gin.Context) {
	id := c.Param("id")
	logs, err := s.controller.GetLogs(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Type", "text/plain")
	c.String(200, logs)
}

func (s *Server) handleRegister(c *gin.Context) {
	var input models.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := s.auth.Register(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (s *Server) handleLogin(c *gin.Context) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := s.auth.Login(credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(200, gin.H{"token": token})
}

func (s *Server) Start(addr string) error {
	log.Printf("Starting server on %s...", addr)
	return s.router.Run(addr)
}
