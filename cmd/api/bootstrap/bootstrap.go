package bootstrap

import (
	"context"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/infrastructure/mongo/photorepository"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/infrastructure/mongo/reciperepository"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/photo"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/photosaver"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/recipesaver"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/recipesearcher"
	"github.com/ariel-rubilar/photography-api/internal/env"
	"github.com/ariel-rubilar/photography-api/internal/projection/photoview/infrastructure/mongo/photoreadrepository"
	"github.com/ariel-rubilar/photography-api/internal/projection/photoview/infrastructure/mongo/photoviewdrepository"
	"github.com/ariel-rubilar/photography-api/internal/projection/photoview/infrastructure/mongo/recipereadrepository"
	savephotoviewsaver "github.com/ariel-rubilar/photography-api/internal/projection/photoview/usecases/photoviewsaver"

	"github.com/ariel-rubilar/photography-api/internal/server"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/imbus"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/mongo"
	webphotorepository "github.com/ariel-rubilar/photography-api/internal/web/infrastructure/mongo/photorepository"
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

	photoReadRepository := photoreadrepository.NewMongoRepository(mongoClient)

	recipeReadRepository := recipereadrepository.NewMongoRepository(mongoClient)

	photoViewRepository := photoviewdrepository.NewMongoRepository(mongoClient)

	photoViewProjector := savephotoviewsaver.NewSavePhotoViewOnPhotoCreated(photoReadRepository, recipeReadRepository, photoViewRepository)

	bus.Subscribe(photo.PhotoCreatedEventType, photoViewProjector)

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
