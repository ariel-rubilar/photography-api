package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, err any) {
		log.Printf("panic: %v", err)
		c.AbortWithStatusJSON(500, gin.H{
			"error": "internal_server_error",
		})
	})
}
