package bot

import (
	"fmt"
	"log"

	"github.com/KLX1899/KiMovieBot/internal/tmdb"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(token, tmdbKey string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true
	log.Printf("✅ Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore non-message updates
			continue
		}

		query := update.Message.Text
		movies, err := tmdb.SearchMovie(tmdbKey, query)
		if err != nil || len(movies) == 0 {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "❌ No results found.")
			bot.Send(msg)
			continue
		}

		movie := movies[0] // Return the first result
		posterURL := fmt.Sprintf("https://image.tmdb.org/t/p/w500%s", movie.PosterPath)
		text := fmt.Sprintf("🎬 *%s*\n⭐ Rating: %.1f\n\n%s", movie.Title, movie.VoteAverage, movie.Overview)

		msg := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileURL(posterURL))
		msg.Caption = text
		msg.ParseMode = "Markdown"
		bot.Send(msg)
	}
}
