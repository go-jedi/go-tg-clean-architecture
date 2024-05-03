package handler

import (
	"context"

	"github.com/go-jedi/go-telegram-clean-architecture/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	ctx         context.Context
	Bot         *tgbotapi.BotAPI
	userService service.UserService
}

func NewHandler(ctx context.Context, bot *tgbotapi.BotAPI, userService service.UserService) *Handler {
	return &Handler{
		ctx:         ctx,
		Bot:         bot,
		userService: userService,
	}
}
