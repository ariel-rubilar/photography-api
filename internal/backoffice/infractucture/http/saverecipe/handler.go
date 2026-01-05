package saverecipe

import (
	"fmt"
	"net/http"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/recipesaver"
	"github.com/gin-gonic/gin"
)

func NewHandler(searcher *recipesaver.Saver) gin.HandlerFunc {
	return func(c *gin.Context) {

		var request recipeDTO

		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("invalid request body: %v", err),
			})
			return
		}

		recipe, err := request.toDomain()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("invalid recipe data: %v", err),
			})
			return
		}

		err = searcher.Save(c.Request.Context(), recipe)

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
