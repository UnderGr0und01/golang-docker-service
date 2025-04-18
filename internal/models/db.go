package models

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(isTest bool) error {
	envFile := ".env"
	if isTest {
		envFile = ".env.test"
	}

	if err := godotenv.Load(envFile); err != nil {
		log.Printf("Warning: Error loading %s file: %v", envFile, err)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		if i < maxRetries-1 {
			time.Sleep(2 * time.Second)
		}
	}
	if err != nil {
		return fmt.Errorf("failed to connect to database after %d retries: %v", maxRetries, err)
	}

	if isTest {
		if err := DB.Exec("DROP TABLE IF EXISTS users CASCADE").Error; err != nil {
			return fmt.Errorf("failed to drop users table: %v", err)
		}
	}

	if err := DB.AutoMigrate(&User{}); err != nil {
		return fmt.Errorf("failed to migrate database: %v", err)
	}

	return nil
}
