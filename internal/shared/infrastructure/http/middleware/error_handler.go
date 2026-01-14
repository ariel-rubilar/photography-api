package middleware

import (
	"net/http"

	domainerror "github.com/ariel-rubilar/photography-api/internal/shared/domain/error"
	sharedhttp "github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http/httperror"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ErrorHandler(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err

		status := http.StatusInternalServerError
		code := "INTERNAL_ERROR"
		msg := "internal server error"

		if he, ok := err.(httperror.Error); ok {
			status = he.StatusCode
			err = he.Err
			msg = he.Message
			code = he.Code
		}

		if de, ok := err.(domainerror.Error); ok {
			code = de.Code()
			msg = de.Error()
		}

		fields := []zap.Field{
			zap.String("error_code", code),
			zap.Int("http_status", status),
			zap.String("path", c.FullPath()),
			zap.String("method", c.Request.Method),
			zap.String("message", msg),
		}

		if status >= 500 {
			logger.Error("request_failed",
				append(fields, zap.Error(err), zap.Stack("stack"))...,
			)
		} else {
			logger.Warn("request_failed",
				append(fields, zap.Error(err))...,
			)
		}

		c.JSON(status, sharedhttp.NewErrorResponse(code, msg))
	}
}
