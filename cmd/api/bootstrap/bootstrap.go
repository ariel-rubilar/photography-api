package bootstrap

import (
	"context"

	"github.com/ariel-rubilar/photography-api/internal/backofice/infractucture/mongo/photorepository"
	"github.com/ariel-rubilar/photography-api/internal/backofice/infractucture/mongo/reciperepository"
	"github.com/ariel-rubilar/photography-api/internal/backofice/usecases/photosaver"
	"github.com/ariel-rubilar/photography-api/internal/backofice/usecases/recipesaver"
	"github.com/ariel-rubilar/photography-api/internal/backofice/usecases/recipesearcher"
	"github.com/ariel-rubilar/photography-api/internal/env"

	"github.com/ariel-rubilar/photography-api/internal/server"
	"github.com/ariel-rubilar/photography-api/internal/shared/infractucture/imbus"
	"github.com/ariel-rubilar/photography-api/internal/shared/infractucture/mongo"
	webphotorepository "github.com/ariel-rubilar/photography-api/internal/web/infractucture/mongo/photorepository"
	"github.com/ariel-rubilar/photography-api/internal/web/usecases/searcher"
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

	webPhotoRepository := webphotorepository.NewMongoRepository(mongoClient)

	photoRepository := photorepository.NewMongoRepository(mongoClient)

	photoSearcherUseCase := searcher.New(webPhotoRepository)

	recipeRepository := reciperepository.NewMongoRepository(mongoClient)

	recipeSearcherUseCase := recipesearcher.New(recipeRepository)

	recipeSaverUseCase := recipesaver.New(recipeRepository)

	photoSaverUseCase := photosaver.New(photoRepository, bus)

	providers := &server.Providers{
		PhotoSearcher:  photoSearcherUseCase,
		RecipeSearcher: recipeSearcherUseCase,
		RecipeSaver:    recipeSaverUseCase,
		PhotoSaver:     photoSaverUseCase,
	}

	s := server.New(providers)

	return s.Start()
}
