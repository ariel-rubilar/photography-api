package live

import (
	"github.com/gin-gonic/gin"
)

func NewHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "alive",
		})
	}
}
