package config

import (
	"os"

	"github.com/pkg/errors"
)

const (
	tokenEnvName = "TOKEN"
)

type TelegramConfig interface {
	TOKEN() string
}

type telegramConfig struct {
	token string
}

func NewTelegramConfig() (TelegramConfig, error) {
	token := os.Getenv(tokenEnvName)
	if len(token) == 0 {
		return nil, errors.New("telegram token not found")
	}

	return &telegramConfig{
		token: token,
	}, nil
}

func (cfg *telegramConfig) TOKEN() string {
	return cfg.token
}
