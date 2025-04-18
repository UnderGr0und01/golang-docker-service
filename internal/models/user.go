package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"unique;not null"`
	Password  string    `json:"-" gorm:"not null;type:text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) CheckPassword(password string) bool {
	log.Printf("DEBUG: Comparing password with hash: %s", u.Password)
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		log.Printf("DEBUG: Password comparison failed: %v", err)
		return false
	}
	log.Printf("DEBUG: Password comparison successful")
	return true
}

func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func InitTestDB() (*gorm.DB, error) {
	// Use test environment variables
	os.Setenv("DB_HOST", "postgres-test")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_NAME", "docker_service_test")

	InitDB(true)

	// Drop and recreate users table
	if err := DB.Migrator().DropTable(&User{}); err != nil {
		return nil, fmt.Errorf("failed to drop users table: %v", err)
	}

	// Auto migrate the schema
	if err := DB.AutoMigrate(&User{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	return DB, nil
}

func (u *User) GenerateToken(secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Username,
		"exp":      time.Now().Add(time.Hour * 24 * 365).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}
