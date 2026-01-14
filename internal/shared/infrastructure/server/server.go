package server

import (
	"context"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type server struct {
	handler http.Handler
	logger  *zap.Logger
}

func New(handler http.Handler, logger *zap.Logger) *server {

	srv := &server{
		handler: handler,
		logger:  logger,
	}

	return srv
}

func (s *server) Start(ctx context.Context) error {

	srv := &http.Server{
		Addr:    ":8080",
		Handler: s.handler,
	}
	s.logger.Info("server started")

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatal("listen: %s\n", zap.Error(err))
		}
	}()

	<-ctx.Done()
	s.logger.Info("shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return srv.Shutdown(shutdownCtx)
}
