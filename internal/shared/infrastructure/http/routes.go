package http

import (
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http/handler/live"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http/handler/ready"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Providers struct {
	DB *mongo.Client
}

func RegisterRoutes(rg *gin.RouterGroup, providers *Providers) {
	rg.GET("/", live.NewHandler())
	rg.GET("/livez", live.NewHandler())
	rg.GET("/readyz", ready.NewHandler(providers.DB))
}
