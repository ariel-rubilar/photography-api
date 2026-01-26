package httpgin

import (
	backofficehttp "github.com/ariel-rubilar/photography-api/internal/backoffice/infrastructure/http"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/photosaver"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/recipesaver"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/recipesearcher"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/uploadurlgetter"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/env"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http/handler"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http/health"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http/middleware"
	webhttp "github.com/ariel-rubilar/photography-api/internal/web/infrastructure/http"
	"github.com/ariel-rubilar/photography-api/internal/web/usecases/searcher"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type Providers struct {
	RecipeSearcher  *recipesearcher.Searcher
	RecipeSaver     *recipesaver.Saver
	PhotoSaver      *photosaver.Saver
	PhotoSearcher   *searcher.Searcher
	DB              *mongo.Client
	Logger          *zap.Logger
	UploadURLGetter *uploadurlgetter.Getter
}

type Config struct {
	Env            env.Env
	GoogleClientID string
}

func NewGinEngine(cfg Config, providers *Providers) *gin.Engine {

	if cfg.Env == env.Development {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	e := gin.New()

	e.NoRoute(handler.NoFound())
	e.NoMethod(handler.NoMethod())

	e.Use(
		middleware.RequestID(),
		middleware.Logger(providers.Logger),
		middleware.Recovery(providers.Logger),
		middleware.ErrorHandler(providers.Logger),
	)

	health.RegisterRoutes(e.Group("/"), &health.Providers{
		DB: providers.DB,
	})

	registerRoutes(e, providers, cfg)

	return e
}

func registerRoutes(r *gin.Engine, providers *Providers, cfg Config) {

	apiVersionGroup := r.Group("/api/v1")

	backofficeGroup := apiVersionGroup.Group("/backoffice")

	backofficehttp.RegisterRoutes(backofficeGroup, &backofficehttp.Providers{
		RecipeSearcher:  providers.RecipeSearcher,
		RecipeSaver:     providers.RecipeSaver,
		PhotoSaver:      providers.PhotoSaver,
		UploadURLGetter: providers.UploadURLGetter,
	}, backofficehttp.Config{
		GoogleClientID: cfg.GoogleClientID,
	})

	webGroup := apiVersionGroup.Group("/web")

	webhttp.RegisterRoutes(webGroup, &webhttp.Providers{
		PhotoSearcher: providers.PhotoSearcher,
	})
}
