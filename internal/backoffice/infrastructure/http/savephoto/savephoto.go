package savephoto

import (
	"net/http"

	sharedhttp "github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/photosaver"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http/httperror"
	"github.com/gin-gonic/gin"
)

func NewHandler(searcher *photosaver.Saver) gin.HandlerFunc {
	return func(c *gin.Context) {

		var request photoDTO

		if err := c.ShouldBind(&request); err != nil {
			c.Error(httperror.WrapBadRequestError(err, httperror.WithMessage("invalid request body")))
			return
		}

		err := searcher.Save(c.Request.Context(), request.ID, request.Title, request.URL, request.RecipeID)

		if err != nil {
			c.Error(httperror.WrapInternalServerError(err))
			return
		}

		c.JSON(http.StatusOK, sharedhttp.NewSuccessResponse("photo saved successfully"))

	}
}
