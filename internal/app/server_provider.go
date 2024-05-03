package app

import (
	"context"
	"log"

	"github.com/go-jedi/go-telegram-clean-architecture/internal/config"
	"github.com/go-jedi/go-telegram-clean-architecture/internal/repository"
	"github.com/go-jedi/go-telegram-clean-architecture/internal/service"
	"github.com/go-jedi/platform_common/pkg/closer"
	"github.com/go-jedi/platform_common/pkg/db"
	"github.com/go-jedi/platform_common/pkg/db/pg"
	"github.com/go-jedi/platform_common/pkg/db/transaction"
)

type serverProvider struct {
	loggerConfig   config.LoggerConfig
	pgConfig       config.PGConfig
	telegramConfig config.TelegramConfig

	dbClient  db.Client
	txManager db.TxManager

	userRepository repository.UserRepository
	userService    service.UserService
}

func newServerProvider() *serverProvider {
	return &serverProvider{}
}

func (s *serverProvider) LoggerConfig() config.LoggerConfig {
	if s.loggerConfig == nil {
		cfg, err := config.NewLoggerConfig()
		if err != nil {
			log.Fatalf("failed to get logger config: %s", err.Error())
		}

		s.loggerConfig = cfg
	}

	return s.loggerConfig
}

func (s *serverProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serverProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serverProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serverProvider) TelegramConfig() config.TelegramConfig {
	if s.telegramConfig == nil {
		cfg, err := config.NewTelegramConfig()
		if err != nil {
			log.Fatalf("failed to get telegram config: %s", err.Error())
		}

		s.telegramConfig = cfg
	}

	return s.telegramConfig
}
