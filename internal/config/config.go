package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramToken string
	TMDBApiKey    string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ No .env file found. Using system environment variables.")
	}

	cfg := Config{
		TelegramToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		TMDBApiKey:    os.Getenv("TMDB_API_KEY"),
	}

	if cfg.TelegramToken == "" || cfg.TMDBApiKey == "" {
		log.Fatal("❌ Missing TELEGRAM_BOT_TOKEN or TMDB_API_KEY in environment.")
	}

	return cfg
}
