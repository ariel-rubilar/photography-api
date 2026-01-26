package recipesaver

import (
	"context"
	"fmt"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/recipe"
	"github.com/ariel-rubilar/photography-api/internal/shared/domain/event"
)

type Saver struct {
	repo recipe.Repository
	bus  event.Bus
}

func New(repo recipe.Repository, bus event.Bus) *Saver {
	return &Saver{
		repo: repo,
		bus:  bus,
	}
}

func (s *Saver) Execute(ctx context.Context, new *recipe.Recipe) error {

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

	if len(recipes) > 0 {
		return fmt.Errorf("recipe with id %s already exists", new.ID())
	}

	err = s.repo.Save(ctx, new)
	if err != nil {
		return err
	}

	return s.bus.Publish(ctx, new.PullEvents())
}
