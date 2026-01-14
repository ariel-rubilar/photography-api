package bootstrap

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ariel-rubilar/photography-api/internal/env"
	"go.uber.org/zap"

	"github.com/ariel-rubilar/photography-api/internal/server"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/imbus"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/mongo"
)

func Run(cfg env.Config, logger *zap.Logger) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-quit
		cancel()
	}()

	mongoClient, err := mongo.NewClient(cfg.MongoURI)

	if err != nil {
		return fmt.Errorf("failed to create mongo client: %w", err)
	}

	defer func() {
		_ = mongoClient.Disconnect(ctx)
	}()

	bus := imbus.New()

	setupProjection(mongoClient, bus)

	webProviders := setupWeb(mongoClient)

	backofficeProviders := setupBackoffice(mongoClient, bus)

	providers := &server.Providers{
		PhotoSearcher:  webProviders.PhotoSearcher,
		RecipeSearcher: backofficeProviders.RecipeSearcher,
		RecipeSaver:    backofficeProviders.RecipeSaver,
		PhotoSaver:     backofficeProviders.PhotoSaver,
		DB:             mongoClient,
		Logger:         logger,
	}

	s := server.New(server.Config{
		Env: cfg.ServerEnv,
	}, providers)

	return s.Start(ctx)
}
