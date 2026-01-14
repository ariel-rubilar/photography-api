package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)

		fields := []zap.Field{
			zap.String("method", c.Request.Method),
			zap.String("path", c.FullPath()),
			zap.String("raw_path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),

			zap.Int("status", c.Writer.Status()),
			zap.Int("response_size", c.Writer.Size()),

			zap.Duration("duration", duration),
			zap.Int64("duration_ms", duration.Milliseconds()),

			zap.String("client_ip", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
			zap.String("host", c.Request.Host),
		}

		if reqID := c.GetHeader("X-Request-ID"); reqID != "" {
			fields = append(fields, zap.String("request_id", reqID))
		}

		switch {
		case c.Writer.Status() >= 500:
			logger.Error("http_request", fields...)
		case c.Writer.Status() >= 400:
			logger.Warn("http_request", fields...)
		default:
			logger.Info("http_request", fields...)
		}

	}
}
