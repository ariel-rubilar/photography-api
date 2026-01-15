package getuploadurl

import (
	"net/http"
	"time"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/uploadurlgetter"
	sharedhttp "github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http"

	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http/httperror"
	"github.com/gin-gonic/gin"
)

type GetUploadURLHandler struct {
	uc *uploadurlgetter.Getter
}

func NewHandler(getter *uploadurlgetter.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request struct {
			ContentType string     `json:"contentType"`
			Extension   string     `json:"extension"`
			TakenAt     *time.Time `json:"takenAt"`
		}

		if err := c.ShouldBind(&request); err != nil {
			c.Error(httperror.WrapBadRequestError(err, httperror.WithMessage("invalid request body")))
			return
		}

		res, err := getter.Execute(
			c.Request.Context(),
			uploadurlgetter.GetUploadURLCommand{
				TakenAt:     request.TakenAt,
				ContentType: request.ContentType,
				Extension:   request.Extension,
			},
		)

		if err != nil {
			c.Error(httperror.WrapInternalServerError(err))

			return
		}

		c.JSON(http.StatusOK, sharedhttp.NewSuccessResponse(Response{
			UploadURL: res.UploadURL,
			ObjectKey: res.ObjectKey,
		}))
	}
}
