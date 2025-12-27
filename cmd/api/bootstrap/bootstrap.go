package bootstrap

import (
	"context"

	"github.com/ariel-rubilar/photography-api/internal/backofice/infractucture/mongo/reciperepository"
	"github.com/ariel-rubilar/photography-api/internal/backofice/usecases/recipesearcher"
	"github.com/ariel-rubilar/photography-api/internal/server"
	"github.com/ariel-rubilar/photography-api/internal/web/infractucture/env"
	"github.com/ariel-rubilar/photography-api/internal/web/infractucture/mongo"
	"github.com/ariel-rubilar/photography-api/internal/web/infractucture/mongo/photorepository"
	"github.com/ariel-rubilar/photography-api/internal/web/usecases/searcher"
)

func Run() error {

	ctx := context.Background()

	cfg, err := env.LoadConfig()

	mongoClient, err := mongo.NewMongoClient(cfg.MongoURI)

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

	photoRepository := photorepository.NewMongoRepository(mongoClient)

	photoSearcherUseCase := searcher.New(photoRepository)

	recipeRepository := reciperepository.NewMongoRepository(mongoClient)

	recipeSearcherUseCase := recipesearcher.New(recipeRepository)

	providers := &server.Providers{
		PhotoSearcher:  photoSearcherUseCase,
		RecipeSearcher: recipeSearcherUseCase,
	}

	s := server.New(providers)

	return s.Start()
}
