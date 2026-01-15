package searcher

import (
	"context"

	"github.com/ariel-rubilar/photography-api/internal/web/photo"
)

type Searcher struct {
	repo photo.Repository
}

func New(repo photo.Repository) *Searcher {
	return &Searcher{
		repo: repo,
	}
}

func (s *Searcher) Execute(ctx context.Context) ([]*photo.Photo, error) {
	return s.repo.Search(ctx)
}
