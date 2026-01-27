package photosearcher

import (
	"context"

	"github.com/ariel-rubilar/photography-api/internal/web/usecases/photoquery"
)

type Searcher struct {
	repo photoquery.Repository
}

func New(repo photoquery.Repository) *Searcher {
	return &Searcher{
		repo: repo,
	}
}

func (s *Searcher) Execute(ctx context.Context) ([]*photoquery.PhotoDTO, error) {
	return s.repo.Search(ctx, photoquery.Criteria{})
}
