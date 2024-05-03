package handler

import (
	"github.com/go-jedi/go-telegram-clean-architecture/internal/command"
	"github.com/go-jedi/go-telegram-clean-architecture/internal/command/about"
	"github.com/go-jedi/go-telegram-clean-architecture/internal/command/help"
	"github.com/go-jedi/go-telegram-clean-architecture/internal/command/start"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handler) HandleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.From.ID, "Я не знаю такой команды")

	switch message.Command() {
	case command.Start:
		msg.ParseMode = "Markdown"
		err := start.Start(h.ctx, h.Bot, h.userService, msg, message.From.ID, message.From.UserName)
		if err != nil {
			return err
		}
	case command.About:
		msg.ParseMode = "Markdown"
		err := about.About(h.Bot, msg, message.From.ID, message.From.UserName)
		if err != nil {
			return err
		}
	case command.Help:
		err := help.Help(h.Bot, msg, message.From.ID, message.From.UserName)
		if err != nil {
			return err
		}
	default:
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true) // удалить клавиатуру
		_, err := h.Bot.Send(msg)
		if err != nil {
			return err
		}
	}

	return nil
}
