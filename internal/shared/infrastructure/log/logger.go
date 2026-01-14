package log

import (
	"github.com/ariel-rubilar/photography-api/internal/env"
	"go.uber.org/zap"
)

type Config struct {
	Env env.Env
}

func New(cfg Config) (*zap.Logger, error) {
	if cfg.Env == env.Development {
		return zap.NewDevelopment()
	}
	return zap.NewProduction()
}
