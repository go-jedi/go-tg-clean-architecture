package app

import (
	"context"

	"github.com/go-jedi/go-telegram-clean-architecture/internal/repository"
	userRepository "github.com/go-jedi/go-telegram-clean-architecture/internal/repository/user"
	"github.com/go-jedi/go-telegram-clean-architecture/internal/service"
	userService "github.com/go-jedi/go-telegram-clean-architecture/internal/service/user"
)

func (s *serverProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepository(s.DBClient(ctx))
	}

	return s.userRepository
}

func (s *serverProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserRepository(ctx),
			s.TxManager(ctx),
			s.validator,
		)
	}

	return s.userService
}
