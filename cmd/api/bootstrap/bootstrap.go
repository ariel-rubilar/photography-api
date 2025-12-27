package bootstrap

import (
	"context"

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

	defer func() {
		_ = mongoClient.Disconnect(ctx)
	}()

	if err != nil {
		return err
	}

	photoRepository := photorepository.NewMongoRepository(mongoClient)

	photoSearcher := searcher.New(photoRepository)

	providers := &server.Providers{
		PhotoSearcher: photoSearcher,
	}

	s := server.New(providers)

	return s.Start()
}
