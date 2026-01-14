package handler

import (
	"errors"
	"net/http"

	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http/httperror"
	"github.com/gin-gonic/gin"
)

func NoFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Error(
			httperror.Wrap(
				errors.New("not found"),
				http.StatusNotFound,
				httperror.WithMessage("route not found"),
				httperror.WithCode("NOT_FOUND"),
			),
		)
	}
}
