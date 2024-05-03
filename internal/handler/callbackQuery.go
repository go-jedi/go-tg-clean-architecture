package handler

import (
	"fmt"

	"github.com/go-jedi/go-telegram-clean-architecture/internal/command/terms"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handler) CallbackQuery(callbackQuery tgbotapi.CallbackQuery) error {
	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, "callbackQuery")

	if callbackQuery.Data == "DG_TERMS" {
		fmt.Println("DG_TERMS")
		err := terms.Terms(h.Bot, msg, callbackQuery.Message.Chat.ID, callbackQuery.Message.Chat.UserName)
		if err != nil {
			return err
		}
	}

	return nil
}
