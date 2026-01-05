package savephoto

import (
	"fmt"
	"net/http"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/photosaver"
	"github.com/gin-gonic/gin"
)

func NewHandler(searcher photosaver.Saver) gin.HandlerFunc {
	return func(c *gin.Context) {

		var request photoDTO

		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("invalid request body: %v", err),
			})
			return
		}

		err := searcher.Save(c.Request.Context(), request.ID, request.Title, request.URL, request.RecipeID)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("failed to save recipe: %v", err),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "recipe saved successfully",
		})

	}
}
