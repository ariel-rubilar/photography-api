package photosaver

import (
	"context"
	"fmt"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/photo"
	domainerror "github.com/ariel-rubilar/photography-api/internal/shared/domain/error"
	"github.com/ariel-rubilar/photography-api/internal/shared/domain/event"
)

type Saver struct {
	repo photo.Repository
	bus  event.Bus
}

func New(repo photo.Repository, bus event.Bus) *Saver {
	return &Saver{
		repo: repo,
		bus:  bus,
	}
}

func (s *Saver) Save(ctx context.Context, id, title, url, recipeID string) error {

	if err := s.ensurePhotoDoNotExists(ctx, id); err != nil {
		return err
	}

	newPhoto, err := photo.Create(id, title, url, recipeID)

	if err != nil {
		return err
	}

	err = s.repo.Save(ctx, newPhoto)
	if err != nil {
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
