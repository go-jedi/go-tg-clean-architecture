package bot

import (
	"github.com/go-jedi/go-telegram-clean-architecture/internal/handler"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	Bot     *tgbotapi.BotAPI
	Handler *handler.Handler
}

func NewBot(bot *tgbotapi.BotAPI, handler *handler.Handler) *Bot {
	return &Bot{
		Bot:     bot,
		Handler: handler,
	}
}
