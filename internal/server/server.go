package server

import (
	"github.com/ariel-rubilar/photography-api/internal/backofice/infractucture/http/saverecipe"
	"github.com/ariel-rubilar/photography-api/internal/backofice/infractucture/http/searchrecipes"
	"github.com/ariel-rubilar/photography-api/internal/backofice/usecases/recipesaver"
	"github.com/ariel-rubilar/photography-api/internal/backofice/usecases/recipesearcher"
	"github.com/ariel-rubilar/photography-api/internal/web/infractucture/http/searchphotos"
	"github.com/ariel-rubilar/photography-api/internal/web/usecases/searcher"
	"github.com/gin-gonic/gin"
)

type Providers struct {
	PhotoSearcher  *searcher.Searcher
	RecipeSearcher *recipesearcher.Searcher
	RecipeSaver    *recipesaver.Saver
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

	backofficeGroup := s.engine.Group("/backoffice")

	backofficeApiGroup := backofficeGroup.Group("/api/v1")

	backofficeApiGroup.GET("/recipes", searchrecipes.NewHandler(s.providers.RecipeSearcher))
	backofficeApiGroup.POST("/recipes", saverecipe.NewHandler(s.providers.RecipeSaver))

	webGroup := s.engine.Group("/web")

	webApiGroup := webGroup.Group("/api/v1")

	webApiGroup.GET("/photos", searchphotos.NewHandler(s.providers.PhotoSearcher))
	return s.engine.Run()
}
