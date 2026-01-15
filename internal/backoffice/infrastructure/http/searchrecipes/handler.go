package searchrecipes

import (
	"net/http"

	sharedhttp "github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/recipesearcher"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http/httperror"
	"github.com/gin-gonic/gin"
)

func NewHandler(searcher *recipesearcher.Searcher) gin.HandlerFunc {
	return func(c *gin.Context) {

		photos, err := searcher.Execute(c.Request.Context())

		if err != nil {
			c.Error(httperror.WrapInternalServerError(err))
			return
		}

		c.JSON(http.StatusOK, sharedhttp.NewSuccessResponse(newSearchRecipesData(photos)))

	}
}
