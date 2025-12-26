package server

import (
	"github.com/ariel-rubilar/photography-api~/internal/web/infractucture/http/searchphotos"
	"github.com/ariel-rubilar/photography-api~/internal/web/usecases/searcher"
	"github.com/gin-gonic/gin"
)

type Providers struct {
	PhotoSearcher *searcher.Searcher
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

	webGroup := s.engine.Group("/web")

	apiGroup := webGroup.Group("/api/v1")

	apiGroup.GET("/photos", searchphotos.NewSearchPhotosHandler(s.providers.PhotoSearcher))
	return s.engine.Run()
}
