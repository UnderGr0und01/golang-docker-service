package config

import (
	"os"
)

type Config struct {
	JWTSecretKey string
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		JWTSecretKey: os.Getenv("JWT_SECRET_KEY"),
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       os.Getenv("DB_PORT"),
		DBUser:       os.Getenv("DB_USER"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		DBName:       os.Getenv("DB_NAME"),
	}

	if cfg.JWTSecretKey == "" || cfg.DBHost == "" || cfg.DBPort == "" ||
		cfg.DBUser == "" || cfg.DBPassword == "" || cfg.DBName == "" {
		return nil, os.ErrInvalid
	}

	return cfg, nil
}
