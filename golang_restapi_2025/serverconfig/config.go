package serverconfig

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort          string
	DatabaseURL         string
	Environment         string
	LogLevel            string
	RedisAddr           string
	RedisPassword       string
	CloudinaryCloudName string
	CloudinaryApiKey    string
	CloudinaryApiSecret string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("Error loading file: %v", err)
	}
	return &Config{
		ServerPort:          getEnv("SERVER_PORT", "8080"),
		DatabaseURL:         getEnv("DATABASE_URL", "postgres"),
		Environment:         getEnv("ENVIRONMENT", "development"),
		LogLevel:            getEnv("LOG_LEVEL", "info"),
		RedisAddr:           getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword:       getEnv("REDIS_PASSWORD", "123456"),
		CloudinaryCloudName: getEnv("CLOUDINARY_CLOUD_NAME", "Root"),
		CloudinaryApiKey:    getEnv("CLOUDINAR_API_KEY", "596377295866749"),
		CloudinaryApiSecret: getEnv("CLOUDINARY_API_SECRET", "DEvzg1p90cEdZACVQX4pCxK6lJY"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
