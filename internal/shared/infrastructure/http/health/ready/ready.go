package ready

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewHandler(db *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 2*time.Second)
		defer cancel()

		if err := db.Ping(ctx, nil); err != nil {
			c.JSON(503, gin.H{
				"status": "not_ready",
				"error":  "db_unreachable",
			})
			return
		}

		c.JSON(200, gin.H{
			"status": "ready",
		})
	}
}
