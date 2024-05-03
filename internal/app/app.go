package app

import (
	"context"

	"github.com/go-jedi/go-telegram-clean-architecture/internal/bot"
	"github.com/go-jedi/go-telegram-clean-architecture/internal/config"
	"github.com/go-jedi/go-telegram-clean-architecture/internal/handler"
	"github.com/go-jedi/go-telegram-clean-architecture/pkg/logger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type App struct {
	serverProvider *serverProvider
	botAPI         *tgbotapi.BotAPI
	telegramBot    *bot.Bot
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	return a.runTelegramServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServerProvider,
		a.initLogger,
		a.initTelegramBotAPI,
		a.initTelegramBot,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServerProvider(_ context.Context) error {
	a.serverProvider = newServerProvider()
	return nil
}

func (a *App) initLogger(_ context.Context) error {
	logger.Init(
		logger.GetCore(
			logger.GetAtomicLevel(a.serverProvider.LoggerConfig().Level()),
		),
	)

	logger.Info("Logger is running")

	return nil
}

func (a *App) initTelegramBotAPI(_ context.Context) error {
	botAPI, err := tgbotapi.NewBotAPI(a.serverProvider.TelegramConfig().TOKEN())
	if err != nil {
		return err
	}

	a.botAPI = botAPI

	return nil
}

func (a *App) initTelegramBot(ctx context.Context) error {
	// Инициализация Services
	userService := a.serverProvider.UserService(ctx)

	h := handler.NewHandler(ctx, a.botAPI, userService)

	telegramBot := bot.NewBot(a.botAPI, h)

	a.telegramBot = telegramBot

	return nil
}

func (a *App) runTelegramServer() error {
	// запуск сервера
	if err := a.telegramBot.Start(); err != nil {
		logger.Fatal("Failed to start telegram bot", zap.Error(err))
		return err
	}

	return nil
}
