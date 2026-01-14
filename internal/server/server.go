package server

import (
	"context"
	"log"
	"net/http"
	"time"

	backofficehttp "github.com/ariel-rubilar/photography-api/internal/backoffice/infrastructure/http"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/photosaver"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/recipesaver"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/recipesearcher"
	webhttp "github.com/ariel-rubilar/photography-api/internal/web/infrastructure/http"
	"github.com/ariel-rubilar/photography-api/internal/web/usecases/searcher"
	"github.com/gin-gonic/gin"
)

type Providers struct {
	RecipeSearcher *recipesearcher.Searcher
	RecipeSaver    *recipesaver.Saver
	PhotoSaver     *photosaver.Saver
	PhotoSearcher  *searcher.Searcher
}

type server struct {
	engine    *gin.Engine
	providers *Providers
}

type Env string

const (
	Development Env = "development"
	Production  Env = "production"
)

type Config struct {
	Env Env
}

func New(cfg Config, providers *Providers) *server {
	e := gin.New()

	if cfg.Env == Development {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	e.Use(gin.Recovery())

	srv := &server{
		engine: e, providers: providers,
	}

	srv.registerRoutes(e)

	return srv
}

func (s *server) Start(ctx context.Context) error {

	srv := &http.Server{
		Addr:    ":8080",
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s", err)
		}
		log.Println("Server is listening on :8080")
	}()

	<-ctx.Done()
	log.Println("shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return srv.Shutdown(shutdownCtx)
}

func (s *server) registerRoutes(r *gin.Engine) {

	apiVersionGroup := r.Group("/api/v1")

	backofficeGroup := apiVersionGroup.Group("/backoffice")

	backofficehttp.RegisterRoutes(backofficeGroup, &backofficehttp.Providers{
		RecipeSearcher: s.providers.RecipeSearcher,
		RecipeSaver:    s.providers.RecipeSaver,
		PhotoSaver:     s.providers.PhotoSaver,
	})

	webGroup := apiVersionGroup.Group("/web")

	webhttp.RegisterRoutes(webGroup, &webhttp.Providers{
		PhotoSearcher: s.providers.PhotoSearcher,
	})
}
