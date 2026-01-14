package searchrecipes

import (
	"net/http"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/recipesearcher"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http/httperror"
	"github.com/gin-gonic/gin"
)

func NewHandler(searcher *recipesearcher.Searcher) gin.HandlerFunc {
	return func(c *gin.Context) {

		photos, err := searcher.Search(c.Request.Context())

		if err != nil {
			c.Error(httperror.Wrap(err, http.StatusInternalServerError, "failed to search recipes"))
			return
		}

		c.JSON(http.StatusOK, newSearchRecipesResponse(photos))

	}
}
