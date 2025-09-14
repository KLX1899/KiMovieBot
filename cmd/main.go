package cmd

import (
	"github.com/KLX1899/KiMovieBot/internal/bot"
	"github.com/KLX1899/KiMovieBot/internal/config"
)

func main() {
	cfg := config.LoadConfig()
	bot.Start(cfg.TelegramToken, cfg.TMDBApiKey)
}
