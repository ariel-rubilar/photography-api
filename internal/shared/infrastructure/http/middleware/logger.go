package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(p gin.LogFormatterParams) string {
		return fmt.Sprintf(
			"%s | %d | %s | %s | %s\n",
			p.TimeStamp.Format(time.RFC3339),
			p.StatusCode,
			p.Method,
			p.Path,
			p.ClientIP,
		)
	})
}
