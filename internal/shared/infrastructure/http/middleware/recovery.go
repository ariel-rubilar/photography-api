package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Recovery(logger *zap.Logger) gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, err any) {
		logger.Error("panic_recovered",
			zap.Any("panic", err),
			zap.Stack("stack"),
		)
		c.AbortWithStatusJSON(500, gin.H{
			"error": "internal_server_error",
		})
	})
}
