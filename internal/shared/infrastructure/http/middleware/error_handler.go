package middleware

import "github.com/gin-gonic/gin"

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err

		switch err.(type) {

		default:
			c.JSON(500, gin.H{"error": "internal_server_error"})
		}
	}
}
