package projector

import (
	"context"
	"fmt"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/photo"
	"github.com/ariel-rubilar/photography-api/internal/projection/photoview/domain/photoread"
	"github.com/ariel-rubilar/photography-api/internal/projection/photoview/domain/photoview"
	"github.com/ariel-rubilar/photography-api/internal/projection/photoview/domain/reciperead"
	"github.com/ariel-rubilar/photography-api/internal/shared/domain/event"
)

type PhotoViewProjector struct {
	photoReader  photoread.Repository
	recipeReader reciperead.Repository
	viewRepo     photoview.Repository
}

func New(
	photoReader photoread.Repository,
	recipeReader reciperead.Repository,
	viewRepo photoview.Repository,
) *PhotoViewProjector {
	return &PhotoViewProjector{
		photoReader:  photoReader,
		recipeReader: recipeReader,
		viewRepo:     viewRepo,
	}
}

func (p *PhotoViewProjector) Handle(ctx context.Context, event event.Event) error {

	e, ok := event.(photo.PhotoCreatedEvent)
	if !ok {
		return fmt.Errorf("invalid event type: %s", event.Type())
	}

	photo, err := p.photoReader.Get(ctx, e.PhotoID())
	if err != nil {
		return err
	}

	recipe, err := p.recipeReader.Get(ctx, e.RecipeID())
	if err != nil {
		return err
	}

	settings := photoview.BuildRecipeSettings(
		recipe.Settings.FilmSimulation,
		recipe.Settings.DynamicRange,
		recipe.Settings.Highlight,
		recipe.Settings.Shadow,
		recipe.Settings.Color,
		recipe.Settings.NoiseReduction,
		recipe.Settings.Sharpening,
		recipe.Settings.Clarity,
		recipe.Settings.GrainEffect,
		recipe.Settings.ColorChromeEffect,
		recipe.Settings.ColorChromeBlue,
		recipe.Settings.WhiteBalance,
		recipe.Settings.Iso,
		recipe.Settings.ExposureCompensation,
	)

	r := photoview.BuildRecipe(
		recipe.ID,
		recipe.Name,
		settings,
		recipe.Link,
	)

	view := photoview.Build(
		photo.ID,
		photo.Title,
		photo.Url,
		r,
	)

	return p.viewRepo.Save(ctx, view)
}
