package bootstrap

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ariel-rubilar/photography-api/internal/env"

	"github.com/ariel-rubilar/photography-api/internal/server"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/imbus"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/mongo"
)

func Run() error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-quit
		cancel()
	}()

	cfg, err := env.LoadConfig()

	if err != nil {
		return err
	}

	mongoClient, err := mongo.NewClient(cfg.MongoURI)

	if err != nil {
		return err
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
	}

	s := server.New(server.Config{
		Env: cfg.ServerEnv,
	}, providers)

	return s.Start(ctx)
}
