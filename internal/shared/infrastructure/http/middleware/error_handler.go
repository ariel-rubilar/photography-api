package middleware

import (
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

		logger.Error("internal_error",
			zap.Error(err),
			zap.Stack("stack"),
		)

		switch err.(type) {

		default:
			c.JSON(500, gin.H{"error": "internal_server_error"})
		}
	}
}
