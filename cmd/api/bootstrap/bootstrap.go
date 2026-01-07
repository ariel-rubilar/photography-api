package bootstrap

import (
	"context"

	"github.com/ariel-rubilar/photography-api/internal/env"

	"github.com/ariel-rubilar/photography-api/internal/server"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/imbus"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/mongo"
)

func Run() error {

	ctx := context.Background()

	cfg, err := env.LoadConfig()

	mongoClient, err := mongo.NewClient(cfg.MongoURI)

	if err != nil {
		return err
	}

	defer func() {
		_ = mongoClient.Disconnect(ctx)
	}()

	err = mongoClient.Ping(ctx, nil)

	if err != nil {
		return err
	}

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

	s := server.New(providers)

	return s.Start()
}
