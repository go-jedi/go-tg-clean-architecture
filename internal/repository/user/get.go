package user

import (
	"context"

	"github.com/go-jedi/go-telegram-clean-architecture/pkg/logger"
	"go.uber.org/zap"
)

func (r *repo) Get(_ context.Context, id int64) (string, error) {
	logger.Info(
		"(REPOSITORY) Get...",
		zap.Int64("id", id),
	)

	return "", nil
}
