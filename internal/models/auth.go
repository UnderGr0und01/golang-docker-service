package models

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type Auth struct {
	db *gorm.DB
}

func NewAuth(db *gorm.DB) *Auth {
	return &Auth{db: db}
}

func (a *Auth) Register(input *UserInput) error {
	var existingUser User
	if err := a.db.Where("username = ?", input.Username).First(&existingUser).Error; err == nil {
		return errors.New("username already exists")
	}

	user := &User{
		Username: input.Username,
	}
	if err := user.HashPassword(input.Password); err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}

	if err := a.db.Create(user).Error; err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	return nil
}

func (a *Auth) Login(username, password string) (string, error) {
	log.Printf("Login attempt for user: %s", username)

	var user User
	if err := a.db.Where("username = ?", username).First(&user).Error; err != nil {
		log.Printf("Error querying user: %v", err)
		return "", errors.New("invalid credentials")
	}

	log.Printf("Found user: %+v", user)
	log.Printf("Comparing passwords: input=%s, stored=%s", password, user.Password)

	if !user.CheckPassword(password) {
		log.Printf("Password check failed")
		return "", errors.New("invalid credentials")
	}

	log.Printf("Password check successful")
	return user.GenerateToken(os.Getenv("JWT_SECRET_KEY"))
}

func (a *Auth) GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24 * 365).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}
