package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DatabaseURL string
}

var AppConfig *Config

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found")
	}
	log.Println("Successfully loaded .env file")
	AppConfig = &Config{
		DatabaseURL: getEnv("DATABASE_URL", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
