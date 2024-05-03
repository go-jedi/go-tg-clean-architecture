package config

import (
	"errors"
	"os"
)

const (
	dsnEnvName = "PG_DSN"
)

var ErrPgDsnNotFound = errors.New("pg dsn not found")

type PGConfig interface {
	DSN() string
}

type pgConfig struct {
	dsn string
}

func NewPGConfig() (PGConfig, error) {
	dsn := os.Getenv(dsnEnvName)
	if len(dsn) == 0 {
		return nil, ErrPgDsnNotFound
	}

	return &pgConfig{
		dsn: dsn,
	}, nil
}

func (cfg *pgConfig) DSN() string {
	return cfg.dsn
}
