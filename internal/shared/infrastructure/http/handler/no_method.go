package handler

import (
	"errors"
	"net/http"

	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http/httperror"
	"github.com/gin-gonic/gin"
)

func NoMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Error(
			httperror.Wrap(
				errors.New("method not allowed"),
				http.StatusMethodNotAllowed,
				httperror.WithMessage("method not allowed"),
				httperror.WithCode("METHOD_NOT_ALLOWED"),
			),
		)
	}
}
