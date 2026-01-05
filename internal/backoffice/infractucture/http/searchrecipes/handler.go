package searchrecipes

import (
	"fmt"
	"net/http"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/recipesearcher"
	"github.com/gin-gonic/gin"
)

func NewHandler(searcher *recipesearcher.Searcher) gin.HandlerFunc {
	return func(c *gin.Context) {

		photos, err := searcher.Search(c.Request.Context())

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("failed to search recipes: %v", err),
			})
			return
		}

		c.JSON(http.StatusOK, newSearchRecipesResponse(photos))

	}
}
