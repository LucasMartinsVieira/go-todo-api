package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	ServerPort string
	APP_ENV    string
}

func LoadConfig() Config {
	_ = godotenv.Load()

	cfg := Config{
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBName:     getEnv("DB_NAME", "todos_db"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
		APP_ENV:    getEnv("APP_ENV", "dev"),
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	log.Printf("⚠️ %s not set, defaulting to %s", key, fallback)
	return fallback
}
