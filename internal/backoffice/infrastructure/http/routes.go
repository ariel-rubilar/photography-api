package http

import (
	"github.com/ariel-rubilar/photography-api/internal/backoffice/infrastructure/http/savephoto"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/infrastructure/http/saverecipe"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/infrastructure/http/searchrecipes"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/photosaver"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/recipesaver"
	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/recipesearcher"
	"github.com/gin-gonic/gin"
)

type Providers struct {
	RecipeSearcher *recipesearcher.Searcher
	RecipeSaver    *recipesaver.Saver
	PhotoSaver     *photosaver.Saver
}

func RegisterRoutes(rg *gin.RouterGroup, providers *Providers) {
	rg.GET("/recipes", searchrecipes.NewHandler(providers.RecipeSearcher))
	rg.POST("/recipes", saverecipe.NewHandler(providers.RecipeSaver))
	rg.POST("/photos", savephoto.NewHandler(providers.PhotoSaver))
}
