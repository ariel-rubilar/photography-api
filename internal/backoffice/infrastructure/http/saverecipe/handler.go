package saverecipe

import (
	"net/http"

	sharedhttp "github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/recipesaver"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http/httperror"
	"github.com/gin-gonic/gin"
)

func NewHandler(searcher *recipesaver.Saver) gin.HandlerFunc {
	return func(c *gin.Context) {

		var request recipeDTO

		if err := c.ShouldBind(&request); err != nil {
			c.Error(httperror.WrapBadRequestError(err, httperror.WithMessage("invalid request body")))

			return
		}

		recipe, err := request.toDomain()

		if err != nil {
			c.Error(httperror.WrapBadRequestError(err, httperror.WithMessage("invalid recipe data")))

			return
		}

		err = searcher.Execute(c.Request.Context(), recipe)

		if err != nil {
			c.Error(httperror.WrapInternalServerError(err))

			return
		}

		c.JSON(http.StatusOK, sharedhttp.NewSuccessResponse("recipe saved successfully"))

	}
}
