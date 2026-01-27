package photosaver

import (
	"context"
	"errors"

	"github.com/ariel-rubilar/photography-api/internal/web/photo"
	"github.com/ariel-rubilar/photography-api/internal/web/usecases/recipequery"
)

type saver struct {
	recipeReader recipequery.Repository
	photoRepo    photo.Repository
}

func New(
	recipeReader recipequery.Repository,
	viewRepo photo.Repository) *saver {
	return &saver{
		recipeReader: recipeReader,
		photoRepo:    viewRepo,
	}
}

func (s *saver) Execute(ctx context.Context, cmd SavePhotoCommand) error {

	exist, err := s.photoRepo.Exists(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if exist {
		return errors.New("photo already exists")
	}

	recipes, err := s.recipeReader.Search(ctx, recipequery.Criteria{
		Filters: recipequery.Filters{
			recipequery.Filter{
				Field: recipequery.FieldID,
				Op:    recipequery.OpEq,
				Value: cmd.RecipeID,
			},
		},
	})

	if err != nil {
		return err
	}

	if len(recipes) == 0 {
		return errors.New("recipe not found")
	}

	if len(recipes) > 1 {
		return errors.New("multiple recipes found with the same ID")
	}

	recipe := recipes[0]

	settings := photo.RecipeSettingsPrimitives{
		FilmSimulation:       recipe.Settings.FilmSimulation,
		DynamicRange:         recipe.Settings.DynamicRange,
		Highlight:            recipe.Settings.Highlight,
		Shadow:               recipe.Settings.Shadow,
		Color:                recipe.Settings.Color,
		NoiseReduction:       recipe.Settings.NoiseReduction,
		Sharpening:           recipe.Settings.Sharpening,
		Clarity:              recipe.Settings.Clarity,
		GrainEffect:          recipe.Settings.GrainEffect,
		ColorChromeEffect:    recipe.Settings.ColorChromeEffect,
		ColorChromeBlue:      recipe.Settings.ColorChromeBlue,
		WhiteBalance:         recipe.Settings.WhiteBalance,
		Iso:                  recipe.Settings.Iso,
		ExposureCompensation: recipe.Settings.ExposureCompensation,
	}

	r := photo.RecipePrimitives{
		ID:       recipe.ID,
		Name:     recipe.Name,
		Settings: settings,
		Link:     recipe.Link,
	}

	view := photo.FromPrimitives(
		photo.PhotoPrimitives{
			ID:     cmd.ID,
			Title:  cmd.Title,
			URL:    cmd.URL,
			Recipe: r,
		},
	)

	return s.photoRepo.Save(ctx, view)
}
