package recipesearcher

import (
	"context"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/recipequery"
)

type Searcher struct {
	repo recipequery.Repository
}

func New(repo recipequery.Repository) *Searcher {
	return &Searcher{
		repo: repo,
	}
}

func (s *Searcher) Execute(ctx context.Context) ([]*recipequery.RecipeDTO, error) {
	return s.repo.Search(ctx, recipequery.Criteria{})
}
