package recipesearcher

import (
	"context"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/recipe"
)

type Searcher struct {
	repo recipe.Repository
}

func New(repo recipe.Repository) *Searcher {
	return &Searcher{
		repo: repo,
	}
}

func (s *Searcher) Execute(ctx context.Context) ([]*recipe.Recipe, error) {
	return s.repo.Search(ctx, recipe.Criteria{})
}
