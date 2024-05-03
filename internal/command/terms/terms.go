package terms

import (
	"fmt"

	"github.com/go-jedi/go-telegram-clean-architecture/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Terms(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig, telegramID int64, userName string) error {
	logger.Info("(COMMAND) Terms...")

	text := "Вызов функции по обработки *inlineButton*"

	msg.Text = text

	fmt.Println("telegramID:", telegramID)
	fmt.Println("userName:", userName)

	_, err := bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
