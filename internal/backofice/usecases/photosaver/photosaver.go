package photosaver

import (
	"context"

	"github.com/ariel-rubilar/photography-api/internal/backofice/photo"
)

type Saver struct {
	repo photo.Repository
}

func New(repo photo.Repository) *Saver {
	return &Saver{
		repo: repo,
	}
}

func (s *Saver) Save(ctx context.Context, id, title, url, recipeID string) error {
	newPhoto, err := photo.Create(id, title, url, recipeID)

	if err != nil {
		return err
	}

	return s.repo.Save(ctx, newPhoto)
}
