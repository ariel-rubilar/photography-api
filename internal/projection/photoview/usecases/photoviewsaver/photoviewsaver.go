package savephotoviewsaver

import (
	"context"

	"github.com/ariel-rubilar/photography-api/internal/projection/photoview/domain/photoread"
	"github.com/ariel-rubilar/photography-api/internal/projection/photoview/domain/photoview"
	"github.com/ariel-rubilar/photography-api/internal/projection/photoview/domain/reciperead"
)

type saver struct {
	photoReader  photoread.Repository
	recipeReader reciperead.Repository
	viewRepo     photoview.Repository
}

func New(photoReader photoread.Repository,
	recipeReader reciperead.Repository,
	viewRepo photoview.Repository) *saver {
	return &saver{
		photoReader:  photoReader,
		recipeReader: recipeReader,
		viewRepo:     viewRepo,
	}
}

func (s *saver) Execute(ctx context.Context, photoID string, recipeID string) error {

	photo, err := s.photoReader.Get(ctx, photoID)
	if err != nil {
		return err
	}

	recipe, err := s.recipeReader.Get(ctx, recipeID)
	if err != nil {
		return err
	}

	settings := photoview.NewRecipeSettings(
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

	r := photoview.NewRecipe(
		recipe.ID,
		recipe.Name,
		settings,
		recipe.Link,
	)

	view := photoview.New(
		photo.ID,
		photo.Title,
		photo.Url,
		r,
	)

	return s.viewRepo.Save(ctx, view)
}
