package bootstrap

import (
	"github.com/ariel-rubilar/photography-api/internal/backoffice/infrastructure/mongo/photorepository"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/infrastructure/mongo/reciperepository"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/photosaver"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/recipesaver"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/recipesearcher"
	"github.com/ariel-rubilar/photography-api/internal/shared/domain/event"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type backofficeProviders struct {
	RecipeSearcher *recipesearcher.Searcher
	RecipeSaver    *recipesaver.Saver
	PhotoSaver     *photosaver.Saver
}

func setupBackoffice(mongoClient *mongo.Client, bus event.Bus) *backofficeProviders {

	photoRepository := photorepository.NewMongoRepository(mongoClient)

	recipeRepository := reciperepository.NewMongoRepository(mongoClient)

	recipeSearcherUseCase := recipesearcher.New(recipeRepository)

	recipeSaverUseCase := recipesaver.New(recipeRepository)

	photoSaverUseCase := photosaver.New(photoRepository, recipeRepository, bus)

	return &backofficeProviders{
		RecipeSearcher: recipeSearcherUseCase,
		RecipeSaver:    recipeSaverUseCase,
		PhotoSaver:     photoSaverUseCase,
	}
}
