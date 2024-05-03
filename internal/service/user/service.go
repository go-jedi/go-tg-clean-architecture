package user

import (
	"github.com/go-jedi/go-telegram-clean-architecture/internal/repository"
	"github.com/go-jedi/go-telegram-clean-architecture/internal/service"
	"github.com/go-jedi/platform_common/pkg/db"
)

type serv struct {
	userRepository repository.UserRepository
	txManager      db.TxManager
}

func NewService(
	userRepository repository.UserRepository,
	txManager db.TxManager,
) service.UserService {
	return &serv{
		userRepository: userRepository,
		txManager:      txManager,
	}
}
