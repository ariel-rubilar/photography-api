package searchphotos

import (
	"fmt"
	"net/http"

	"github.com/ariel-rubilar/photography-api~/internal/web/usecases/searcher"
	"github.com/gin-gonic/gin"
)

func NewSearchPhotosHandler(searcher *searcher.Searcher) gin.HandlerFunc {
	return func(c *gin.Context) {

		photos, err := searcher.Search(c.Request.Context())

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("failed to search photos: %v", err),
			})
			return
		}

		c.JSON(http.StatusOK, NewSearchPhotosResponse(photos))

	}
}
