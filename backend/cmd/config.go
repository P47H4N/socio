package cmd

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	DBHost       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBPort       string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	port := os.Getenv("PORT")
	dbHost := os.Getenv("DB_HOST")
	if port == "" {
		port = "8080"
	}
	if dbHost == "" {
		dbHost = "localhost"
	}
	return &Config{
		Port:       port,
		DBHost:     dbHost,
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
	}, nil
}