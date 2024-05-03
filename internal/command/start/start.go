package start

import (
	"context"
	"fmt"

	"github.com/go-jedi/go-telegram-clean-architecture/internal/service"
	"github.com/go-jedi/go-telegram-clean-architecture/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(ctx context.Context, bot *tgbotapi.BotAPI, service service.UserService, msg tgbotapi.MessageConfig, telegramID int64, userName string) error {
	logger.Info("(COMMAND) Start...")

	text := "Вызов функции по /start команде"

	msg.Text = text

	fmt.Println("telegramID:", telegramID)
	fmt.Println("userName:", userName)

	result, err := service.Get(ctx, 1)
	if err != nil {
		return err
	}
	fmt.Println("result:", result)

	_, err = bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
