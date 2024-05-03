package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) error {
	for update := range updates {
		if update.CallbackQuery != nil {
			err := b.Handler.CallbackQuery(*update.CallbackQuery)
			if err != nil {
				return err
			}
			continue
		}

		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() { // если возвращает true, то это команда, а если false, то нет
			err := b.Handler.HandleCommand(update.Message)
			if err != nil {
				return err
			}
			continue
		}

		b.Handler.HandleMessage(update.Message)
	}

	return nil
}
