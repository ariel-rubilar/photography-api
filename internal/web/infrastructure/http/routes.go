package http

import (
	"github.com/ariel-rubilar/photography-api/internal/web/infrastructure/http/searchphotos"
	"github.com/ariel-rubilar/photography-api/internal/web/usecases/photosearcher"
	"github.com/gin-gonic/gin"
)

type Providers struct {
	PhotoSearcher *photosearcher.Searcher
}

func RegisterRoutes(rg *gin.RouterGroup, providers *Providers) {
	rg.GET("/photos", searchphotos.NewHandler(providers.PhotoSearcher))
}
