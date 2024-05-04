package user

import (
	"context"
	"time"

	"github.com/go-jedi/go-telegram-clean-architecture/internal/model/user"
	"github.com/go-jedi/go-telegram-clean-architecture/pkg/logger"
	"go.uber.org/zap"
)

func (s *serv) Get(ctx context.Context, id int64) (string, error) {
	logger.Info(
		"(SERVICE) Get...",
		zap.Int64("id", id),
	)

	dto := user.User{
		ID:        1,
		Username:  "username",
		IsBan:     false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := s.validator.Struct(dto)
	if err != nil {
		return "", err
	}

	result, err := s.userRepository.Get(ctx, id)
	if err != nil {
		return "", err
	}

	return result, nil
}
