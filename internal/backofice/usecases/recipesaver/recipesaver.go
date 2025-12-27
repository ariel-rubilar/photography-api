package recipesaver

import (
	"context"
	"fmt"

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

func (s *Saver) Save(ctx context.Context, new *recipe.Recipe) error {

	recipes, err := s.repo.Search(ctx, recipe.Criteria{
		Filters: recipe.Filters{
			{
				Field: recipe.FieldID,
				Op:    recipe.OpEq,
				Value: new.ID(),
			},
		},
	})

	if err != nil {
		return err
	}

	fmt.Println("Found recipes:", len(recipes))

	if len(recipes) > 0 {
		return fmt.Errorf("recipe with id %s already exists", new.ID())
	}

	return s.repo.Save(ctx, new)
}
