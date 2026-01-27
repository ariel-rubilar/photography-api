package recipesaver

import (
	"context"
	"fmt"

	bcrecipe "github.com/ariel-rubilar/photography-api/internal/backoffice/recipe"
	"github.com/ariel-rubilar/photography-api/internal/shared/domain/event"
	"github.com/ariel-rubilar/photography-api/internal/web/recipe"
)

type handler struct {
	saver *Saver
}

func NewSaveRecipeOnBackofficeRecipeCreated(
	repository recipe.Repository,
) *handler {
	return &handler{
		saver: New(repository),
	}
}

func (p *handler) Handle(ctx context.Context, event event.Event) error {

	e, ok := event.(bcrecipe.RecipeCreatedEvent)
	if !ok {
		return fmt.Errorf("invalid event type: %s", event.Type())
	}

	cmd := SaveRecipeCommand{
		ID:   e.RecipeID(),
		Name: e.Name(),
		Settings: SaveRecipeSettingsCommand{
			FilmSimulation:       e.Settings().FilmSimulation,
			DynamicRange:         e.Settings().DynamicRange,
			Highlight:            e.Settings().Highlight,
			Shadow:               e.Settings().Shadow,
			Color:                e.Settings().Color,
			NoiseReduction:       e.Settings().NoiseReduction,
			Sharpening:           e.Settings().Sharpening,
			Clarity:              e.Settings().Clarity,
			GrainEffect:          e.Settings().GrainEffect,
			ColorChromeEffect:    e.Settings().ColorChromeEffect,
			ColorChromeBlue:      e.Settings().ColorChromeBlue,
			WhiteBalance:         e.Settings().WhiteBalance,
			Iso:                  e.Settings().Iso,
			ExposureCompensation: e.Settings().ExposureCompensation,
		},
		Link: e.Link(),
	}

	return p.saver.Execute(ctx, cmd)
}
