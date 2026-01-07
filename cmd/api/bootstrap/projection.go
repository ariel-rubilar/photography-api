package bootstrap

import (
	"github.com/ariel-rubilar/photography-api/internal/backoffice/photo"
	"github.com/ariel-rubilar/photography-api/internal/projection/photoview/infrastructure/mongo/photoreadrepository"
	"github.com/ariel-rubilar/photography-api/internal/projection/photoview/infrastructure/mongo/photoviewdrepository"
	"github.com/ariel-rubilar/photography-api/internal/projection/photoview/infrastructure/mongo/recipereadrepository"
	savephotoviewsaver "github.com/ariel-rubilar/photography-api/internal/projection/photoview/usecases/photoviewsaver"
	"github.com/ariel-rubilar/photography-api/internal/shared/domain/event"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func setupProjection(mongoClient *mongo.Client, bus event.Bus) {

	photoReadRepository := photoreadrepository.NewMongoRepository(mongoClient)

	recipeReadRepository := recipereadrepository.NewMongoRepository(mongoClient)

	photoViewRepository := photoviewdrepository.NewMongoRepository(mongoClient)

	photoViewProjector := savephotoviewsaver.NewSavePhotoViewOnPhotoCreated(photoReadRepository, recipeReadRepository, photoViewRepository)

	bus.Subscribe(photo.PhotoCreatedEventType, photoViewProjector)

}
