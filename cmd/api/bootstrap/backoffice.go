package bootstrap

import (
	"github.com/ariel-rubilar/photography-api/internal/backoffice/infrastructure/mongo/photorepository"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/infrastructure/mongo/reciperepository"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/infrastructure/r2"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/photosaver"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/recipesaver"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/recipesearcher"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/uploadurlgetter"
	"github.com/ariel-rubilar/photography-api/internal/shared/aplication/clock"
	"github.com/ariel-rubilar/photography-api/internal/shared/domain/event"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Config struct {
	PublicBaseURL string
	BucketName    string
}

type backofficeProviders struct {
	RecipeSearcher  *recipesearcher.Searcher
	RecipeSaver     *recipesaver.Saver
	PhotoSaver      *photosaver.Saver
	UploadURLGetter *uploadurlgetter.Getter
}

func setupBackoffice(
	cfg Config,
	mongoClient *mongo.Client,
	r2Client *s3.PresignClient,
	clock clock.Clock,
	bus event.Bus,
) *backofficeProviders {

	photoRepository := photorepository.NewMongoRepository(mongoClient)

	recipeRepository := reciperepository.NewMongoRepository(mongoClient)

	recipeSearcherUseCase := recipesearcher.New(recipeRepository)

	recipeSaverUseCase := recipesaver.New(recipeRepository, bus)

	photoSaverUseCase := photosaver.New(photoRepository, recipeRepository, bus)

	urlSigner := r2.NewSigner(cfg.BucketName, r2Client)

	uploadURLGetterUseCase := uploadurlgetter.New(cfg.PublicBaseURL, urlSigner, clock)

	return &backofficeProviders{
		RecipeSearcher:  recipeSearcherUseCase,
		RecipeSaver:     recipeSaverUseCase,
		PhotoSaver:      photoSaverUseCase,
		UploadURLGetter: uploadURLGetterUseCase,
	}
}
