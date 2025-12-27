package bootstrap

import (
	"context"
	"fmt"
	"os"

	"github.com/ariel-rubilar/photography-api/internal/server"
	"github.com/ariel-rubilar/photography-api/internal/web/infractucture/mongo"
	"github.com/ariel-rubilar/photography-api/internal/web/infractucture/mongo/photorepository"
	"github.com/ariel-rubilar/photography-api/internal/web/usecases/searcher"
	"github.com/joho/godotenv"
)

func Run() error {

	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	ctx := context.Background()

	uri := os.Getenv("MONGO_URI")

	if uri == "" {
		return fmt.Errorf("MONGO_URI environment variable is not set")
	}

	mongoClient, err := mongo.NewMongoClient(uri)

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
