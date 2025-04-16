package handlers

import (
	"log"
	"net/http"

	"docker-service/internal/config"
	"docker-service/internal/middleware"
	"docker-service/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
	db *gorm.DB
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		db: config.DB,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	user := models.User{
		Username: input.Username,
	}

	if err := user.HashPassword(input.Password); err != nil {
		log.Printf("Error hashing password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not process password"})
		return
	}

	if err := h.db.Create(&user).Error; err != nil {
		log.Printf("Error creating user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	log.Printf("DEBUG: Login attempt for user: %s", input.Username)

	var user models.User
	if err := h.db.Where("username = ?", input.Username).First(&user).Error; err != nil {
		log.Printf("DEBUG: User not found in database: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	log.Printf("DEBUG: Found user in database: %+v", user)

	if !user.CheckPassword(input.Password) {
		log.Printf("DEBUG: Password check failed for user: %s", input.Username)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := middleware.GenerateToken(user.Username)
	if err != nil {
		log.Printf("Error generating token for user %s: %v", user.Username, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate authentication token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
