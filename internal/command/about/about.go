package about

import (
	"fmt"

	"github.com/go-jedi/go-telegram-clean-architecture/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func About(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, telegramID int64, userName string) error {
	logger.Info("(COMMAND) About...")

	text := "Команда About"

	msg.Text = text

	fmt.Println("telegramID:", telegramID)
	fmt.Println("userName:", userName)

	_, err := bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
