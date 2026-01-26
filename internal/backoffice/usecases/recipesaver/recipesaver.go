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

func (s *Saver) Execute(ctx context.Context, cmd SaveRecipeCommand) error {

	settings := recipe.RecipeSettingsFromPrimitives(recipe.RecipeSettingsPrimitives{
		FilmSimulation:       cmd.Settings.FilmSimulation,
		DynamicRange:         cmd.Settings.DynamicRange,
		Highlight:            cmd.Settings.Highlight,
		Shadow:               cmd.Settings.Shadow,
		Color:                cmd.Settings.Color,
		NoiseReduction:       cmd.Settings.NoiseReduction,
		Sharpening:           cmd.Settings.Sharpening,
		Clarity:              cmd.Settings.Clarity,
		GrainEffect:          cmd.Settings.GrainEffect,
		ColorChromeEffect:    cmd.Settings.ColorChromeEffect,
		ColorChromeBlue:      cmd.Settings.ColorChromeBlue,
		WhiteBalance:         cmd.Settings.WhiteBalance,
		Iso:                  cmd.Settings.Iso,
		ExposureCompensation: cmd.Settings.ExposureCompensation,
	})

	new, err := recipe.Create(
		cmd.ID,
		cmd.Name,
		settings,
		cmd.Link,
	)

	if err != nil {
		return err
	}

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
