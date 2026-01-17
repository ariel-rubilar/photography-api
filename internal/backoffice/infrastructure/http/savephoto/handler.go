package savephoto

import (
	"errors"
	"net/http"

	"github.com/ariel-rubilar/photography-api/internal/shared/domain/domainerror"
	sharedhttp "github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/photosaver"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http/httperror"
	"github.com/gin-gonic/gin"
)

func NewHandler(searcher *photosaver.Saver) gin.HandlerFunc {
	return func(c *gin.Context) {

		var request PhotoDTO

		if err := c.ShouldBind(&request); err != nil {
			c.Error(httperror.WrapBadRequestError(err, httperror.WithMessage("invalid request body")))
			return
		}

		err := searcher.Execute(
			c.Request.Context(),
			request.ID,
			request.Title,
			request.URL,
			request.RecipeID,
		)

		if err != nil {
			var conflictError domainerror.Conflict

			if errors.As(err, &conflictError) {
				c.Error(httperror.Wrap(conflictError, http.StatusConflict))
				return
			}

			var notFoundError domainerror.NotFound

			if errors.As(err, &notFoundError) {
				c.Error(httperror.Wrap(notFoundError, http.StatusNotFound))
				return
			}

			c.Error(httperror.WrapInternalServerError(err))
			return
		}

		c.JSON(http.StatusCreated, sharedhttp.NewSuccessResponse("photo saved successfully"))

	}
}
