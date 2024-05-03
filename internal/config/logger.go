package config

import (
	"os"

	"github.com/pkg/errors"
)

const (
	logLevelEnvName = "LOG_LEVEL"
)

type LoggerConfig interface {
	Level() string
}

type loggerConfig struct {
	level string
}

func NewLoggerConfig() (LoggerConfig, error) {
	level := os.Getenv(logLevelEnvName)
	if len(level) == 0 {
		return nil, errors.New("logger level not found")
	}

	return &loggerConfig{
		level: level,
	}, nil
}

func (cfg *loggerConfig) Level() string {
	return cfg.level
}
