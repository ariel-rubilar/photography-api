package bootstrap

import (
	bcphoto "github.com/ariel-rubilar/photography-api/internal/backoffice/photo"
	bcrecipe "github.com/ariel-rubilar/photography-api/internal/backoffice/recipe"
	"github.com/ariel-rubilar/photography-api/internal/shared/domain/event"
	photorepository "github.com/ariel-rubilar/photography-api/internal/web/infrastructure/mongo/photorepository"
	"github.com/ariel-rubilar/photography-api/internal/web/infrastructure/mongo/reciperepository"
	"github.com/ariel-rubilar/photography-api/internal/web/usecases/photosaver"
	"github.com/ariel-rubilar/photography-api/internal/web/usecases/recipesaver"
	"github.com/ariel-rubilar/photography-api/internal/web/usecases/searcher"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type webProviders struct {
	PhotoSearcher *searcher.Searcher
}

func setupWeb(mongoClient *mongo.Client, bus event.Bus) *webProviders {

	webPhotoRepository := photorepository.NewMongoRepository(mongoClient)

	photoSearcherUseCase := searcher.New(webPhotoRepository)

	webRecipeRepository := reciperepository.NewMongoRepository(mongoClient)

	webrecipesaver := recipesaver.NewSaveRecipeOnBackofficeRecipeCreated(webRecipeRepository)

	webphotosaver := photosaver.NewSaveBackofficePhotoViewOnPhotoCreated(webRecipeRepository, webPhotoRepository)

	bus.Subscribe(bcrecipe.RecipeCreatedEventType, webrecipesaver)
	bus.Subscribe(bcphoto.PhotoCreatedEventType, webphotosaver)

	return &webProviders{
		PhotoSearcher: photoSearcherUseCase,
	}
}
