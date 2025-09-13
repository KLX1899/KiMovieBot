package cmd

import (
	"github.com/klx1899/KiMovieBot/internal/bot"
	"github.com/klx1899/KiMovieBot/internal/config"
)

func main() {
	cfg := config.LoadConfig()
	bot.Start(cfg.TelegramToken, cfg.TMDBApiKey)
}
