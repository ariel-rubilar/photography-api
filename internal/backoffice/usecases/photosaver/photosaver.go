package photosaver

import (
	"context"

	"github.com/ariel-rubilar/photography-api/internal/backoffice/photo"
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
