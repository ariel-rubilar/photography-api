package savephotoviewsaver

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
	saver *saver
}

func NewSavePhotoViewOnPhotoCreated(
	photoReader photoread.Repository,
	recipeReader reciperead.Repository,
	viewRepo photoview.Repository,
) *PhotoViewProjector {
	return &PhotoViewProjector{
		saver: New(photoReader, recipeReader, viewRepo),
	}
}

func (p *PhotoViewProjector) Handle(ctx context.Context, event event.Event) error {

	e, ok := event.(photo.PhotoCreatedEvent)
	if !ok {
		return fmt.Errorf("invalid event type: %s", event.Type())
	}

	return p.saver.Execute(ctx, e.PhotoID(), e.RecipeID())
}
