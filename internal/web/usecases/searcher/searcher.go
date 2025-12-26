package searcher

import (
	"context"

	"github.com/ariel-rubilar/photography-api~/internal/web/domain"
)

type Searcher struct {
	repo domain.Repository
}

func New(repo domain.Repository) *Searcher {
	return &Searcher{
		repo: repo,
	}
}

func (s *Searcher) Search(ctx context.Context) ([]*domain.Photo, error) {
	return s.repo.Search(ctx)
}
