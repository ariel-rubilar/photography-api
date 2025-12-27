package recipesaver

import (
	"context"

	"github.com/ariel-rubilar/photography-api/internal/backofice/recipe"
)

type Saver struct {
	repo recipe.Repository
}

func New(repo recipe.Repository) *Saver {
	return &Saver{
		repo: repo,
	}
}

func (s *Saver) Save(ctx context.Context, recipe *recipe.Recipe) error {
	return s.repo.Save(ctx, recipe)
}
