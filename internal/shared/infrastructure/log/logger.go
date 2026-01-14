package log

import (
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/env"
	"go.uber.org/zap"
)

type Config struct {
	Env env.Env
}

func New(cfg Config) (*zap.Logger, error) {
	zapCfg := zap.NewProductionConfig()
	if cfg.Env == env.Development {
		zapCfg = zap.NewDevelopmentConfig()
	}

	zapCfg.DisableStacktrace = true

	return zapCfg.Build()
}
