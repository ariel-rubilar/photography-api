package photosaver

import (
	"context"
	"fmt"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/photo"
	"github.com/ariel-rubilar/photography-api/internal/shared/domain/domainerror"
	"github.com/ariel-rubilar/photography-api/internal/shared/domain/event"
)

type Saver struct {
	repo       photo.Repository
	bus        event.Bus
	recipeRepo RecipeReadRepository
}

func New(repo photo.Repository, recipeRepo RecipeReadRepository, bus event.Bus) *Saver {
	return &Saver{
		repo:       repo,
		bus:        bus,
		recipeRepo: recipeRepo,
	}
}

func (s *Saver) Execute(ctx context.Context, cmd SavePhotoCommand) error {

	if err := s.ensurePhotoDoNotExists(ctx, cmd.ID); err != nil {
		return err
	}

	if err := s.ensureRecipeExists(ctx, cmd.RecipeID); err != nil {
		return err
	}

	newPhoto, err := photo.Create(cmd.ID, cmd.Title, cmd.URL, cmd.RecipeID)

	if err != nil {
		return err
	}

	if err = s.repo.Save(ctx, newPhoto); err != nil {
		return err
	}

	return s.bus.Publish(ctx, newPhoto.PullEvents())
}

func (s *Saver) ensurePhotoDoNotExists(ctx context.Context, id string) error {
	photos, err := s.repo.Search(ctx, photo.Criteria{
		Filters: photo.Filters{
			{
				Field: photo.FieldID,
				Op:    photo.OpEq,
				Value: id,
			},
		},
	})

	if err != nil {
		return err
	}

	if len(photos) > 0 {
		return domainerror.Conflict{
			Reason: fmt.Sprintf("%s already exists", id),
		}
	}

	return nil
}

func (s *Saver) ensureRecipeExists(ctx context.Context, id string) error {
	ok, err := s.recipeRepo.Exists(ctx, id)

	if err != nil {
		return err
	}

	if !ok {
		return domainerror.NotFound{
			Entity: "recipe",
		}
	}

	return nil
}
