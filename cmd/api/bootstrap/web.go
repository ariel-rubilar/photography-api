package bootstrap

import (
	photorepository "github.com/ariel-rubilar/photography-api/internal/web/infrastructure/mongo/photorepository"
	"github.com/ariel-rubilar/photography-api/internal/web/usecases/searcher"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type webProviders struct {
	PhotoSearcher *searcher.Searcher
}

func setupWeb(mongoClient *mongo.Client) *webProviders {

	webPhotoRepository := photorepository.NewMongoRepository(mongoClient)

	photoSearcherUseCase := searcher.New(webPhotoRepository)

	return &webProviders{
		PhotoSearcher: photoSearcherUseCase,
	}
}
