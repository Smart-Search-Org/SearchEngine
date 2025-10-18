package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Port string
	}
	Database struct {
		DSN string
	}

	OpenAI struct {
		APIKey string
	}
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found, using environment variables only")
	}

	var cfg Config
	cfg.Server.Port = os.Getenv("SERVER_PORT")
	cfg.Database.DSN = os.Getenv("DATABASE_DSN")
	cfg.OpenAI.APIKey = os.Getenv("OPENAI_APIKEY")

	if cfg.Server.Port == "" {
		cfg.Server.Port = "8080"
	}

	return cfg
}
