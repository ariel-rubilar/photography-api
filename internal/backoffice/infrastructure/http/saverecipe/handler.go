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

		err := searcher.Execute(c.Request.Context(), recipesaver.SaveRecipeCommand{
			ID:   request.ID,
			Name: request.Name,
			Settings: recipesaver.SaveRecipeSettingsCommand{
				FilmSimulation:       request.Settings.FilmSimulation,
				DynamicRange:         request.Settings.DynamicRange,
				Highlight:            request.Settings.Highlight,
				Shadow:               request.Settings.Shadow,
				Color:                request.Settings.Color,
				NoiseReduction:       request.Settings.NoiseReduction,
				Sharpening:           request.Settings.Sharpening,
				Clarity:              request.Settings.Clarity,
				GrainEffect:          request.Settings.GrainEffect,
				ColorChromeEffect:    request.Settings.ColorChromeEffect,
				ColorChromeBlue:      request.Settings.ColorChromeBlue,
				WhiteBalance:         request.Settings.WhiteBalance,
				Iso:                  request.Settings.Iso,
				ExposureCompensation: request.Settings.ExposureCompensation,
			},
			Link: request.Link,
		})

		if err != nil {
			c.Error(httperror.WrapInternalServerError(err))

			return
		}

		c.JSON(http.StatusOK, sharedhttp.NewSuccessResponse("recipe saved successfully"))

	}
}
