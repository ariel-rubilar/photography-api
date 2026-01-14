package main

import (
	"fmt"

	"github.com/ariel-rubilar/photography-api/cmd/api/bootstrap"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/env"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/log"
)

func main() {

	cfg, err := env.LoadConfig()

	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	logger, err := log.New(log.Config{
		Env: cfg.ServerEnv,
	})

	if err != nil {
		panic(fmt.Errorf("failed to create logger: %w", err))
	}

	defer logger.Sync()

	if err := bootstrap.Run(*cfg, logger); err != nil {
		panic(fmt.Errorf("failed to run application: %w", err))
	}
}
