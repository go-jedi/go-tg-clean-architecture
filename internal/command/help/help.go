package help

import (
	"fmt"

	"github.com/go-jedi/go-telegram-clean-architecture/internal/keyboard"
	"github.com/go-jedi/go-telegram-clean-architecture/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Help(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, telegramID int64, userName string) error {
	logger.Info("(COMMAND) Help...")

	text := "Ты ввёл команду /help"

	fmt.Println("telegramID:", telegramID)
	fmt.Println("userName:", userName)

	msg.Text = text

	msg.ReplyMarkup = keyboard.GenKeyboardInlineForList("2", "2")

	_, err := bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
