package user

import (
	"github.com/go-jedi/go-telegram-clean-architecture/internal/repository"
	"github.com/go-jedi/go-telegram-clean-architecture/internal/service"
	"github.com/go-jedi/platform_common/pkg/db"
	"github.com/go-playground/validator/v10"
)

type serv struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
	validator      *validator.Validate
}

func NewService(
	userRepository repository.UserRepository,
	txManager db.TxManager,
	validator *validator.Validate,
) service.UserService {
	return &serv{
		userRepository: userRepository,
		txManager:      txManager,
		validator:      validator,
	}
}
