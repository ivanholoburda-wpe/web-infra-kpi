package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbName string
	MqUrl  string
}

func InitConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("Info: .env file not found, reading configuration from environment variables.")
	}

	return &Config{
		DbHost: getEnv("DB_HOST", "localhost"),
		DbPort: getEnv("DB_PORT", "5432"),
		DbUser: getEnv("DB_USER", "user"),
		DbPass: getEnv("DB_PASS", "password"),
		DbName: getEnv("DB_NAME", "sitemonitor"),
		MqUrl:  getEnv("MQ_URL", "amqp://guest:guest@localhost:5672/"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	log.Printf("Warning: Environment variable %s not set. Using default value: %s", key, fallback)
	return fallback
}
