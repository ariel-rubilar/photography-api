package photosaver

import (
	"context"
	"fmt"

	bcphoto "github.com/ariel-rubilar/photography-api/internal/backoffice/photo"
	"github.com/ariel-rubilar/photography-api/internal/shared/domain/event"
	"github.com/ariel-rubilar/photography-api/internal/web/photo"
	"github.com/ariel-rubilar/photography-api/internal/web/usecases/recipequery"
)

type PhotoViewProjector struct {
	saver *saver
}

func NewSaveBackofficePhotoViewOnPhotoCreated(
	recipeReader recipequery.Repository,
	photoRepo photo.Repository,
) *PhotoViewProjector {
	return &PhotoViewProjector{
		saver: New(recipeReader, photoRepo),
	}
}

func (p *PhotoViewProjector) Handle(ctx context.Context, event event.Event) error {

	e, ok := event.(bcphoto.PhotoCreatedEvent)
	if !ok {
		return fmt.Errorf("invalid event type: %s", event.Type())
	}

	cmd := SavePhotoCommand{
		ID:       e.PhotoID(),
		Title:    e.Title(),
		URL:      e.URL(),
		RecipeID: e.PhotoID(),
	}

	return p.saver.Execute(ctx, cmd)
}
