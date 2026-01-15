package bootstrap

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/infrastructure/r2"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/env"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/http/httpgin"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/imbus"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/mongo"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/realclock"
	"github.com/ariel-rubilar/photography-api/internal/shared/infrastructure/server"
)

func Run(cfg env.Config, logger *zap.Logger) error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		<-quit
		cancel()
	}()

	mongoClient, err := mongo.NewClient(cfg.MongoURI)

	if err != nil {
		return fmt.Errorf("failed to create mongo client: %w", err)
	}

	defer func() {
		_ = mongoClient.Disconnect(ctx)
	}()

	bus := imbus.New()

	setupProjection(mongoClient, bus)

	webProviders := setupWeb(mongoClient)

	realClock := realclock.RealClock{}

	r2Client, err := r2.NewClient(ctx, r2.Config{
		AccessKeyID:     cfg.R2.AccessKeyID,
		SecretAccessKey: cfg.R2.SecretAccessKey,
		AccountID:       cfg.R2.AccountID,
	})

	if err != nil {
		return err
	}

	backofficeProviders := setupBackoffice(Config{
		PublicBaseURL: cfg.R2.PublicBaseURL,
		BucketName:    cfg.R2.BucketName,
	}, mongoClient, r2Client, realClock, bus)

	providers := &httpgin.Providers{
		PhotoSearcher:   webProviders.PhotoSearcher,
		RecipeSearcher:  backofficeProviders.RecipeSearcher,
		RecipeSaver:     backofficeProviders.RecipeSaver,
		PhotoSaver:      backofficeProviders.PhotoSaver,
		DB:              mongoClient,
		Logger:          logger,
		UploadURLGetter: backofficeProviders.UploadURLGetter,
	}

	handler := httpgin.NewGinEngine(httpgin.Config{
		Env: cfg.ServerEnv,
	}, providers)

	s := server.New(handler, logger)

	return s.Start(ctx)
}
