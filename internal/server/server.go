package server

import (
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

func New(providers *Providers) *server {
	e := gin.Default()
	return &server{
		engine: e, providers: providers,
	}
}

func (s *server) Start() error {

	s.engine.Use(gin.Recovery())

	s.registerRoutes(s.engine)

	return s.engine.Run()
}

func (s *server) registerRoutes(r *gin.Engine) {

	apiVersionGroup := s.engine.Group("/api/v1")

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
