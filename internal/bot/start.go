package bot

import "github.com/go-jedi/go-telegram-clean-architecture/pkg/logger"

func (b *Bot) Start() error {
	updates, err := b.initUpdatesChannel()
	if err != nil {
		return nil
	}

	logger.Info("Telegram bot is running")

	err = b.handleUpdates(updates)
	if err != nil {
		return err
	}

	return nil
}
