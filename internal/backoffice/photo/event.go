package photo

import "github.com/ariel-rubilar/photography-api/internal/shared/domain/event"

var (
	PhotoCreatedEventType event.Type = "backoffice.events.photo.created"
)

type PhotoCreatedEvent struct {
	event.BaseEvent
	id       string
	recipeID string
}

func newPhotoCreatedEvent(id, recipeID string) PhotoCreatedEvent {
	return PhotoCreatedEvent{
		id:        id,
		recipeID:  recipeID,
		BaseEvent: event.NewBaseEvent(id),
	}
}

func (e PhotoCreatedEvent) Type() event.Type {
	return PhotoCreatedEventType
}

func (e PhotoCreatedEvent) RecipeID() string {
	return e.recipeID
}

func (e PhotoCreatedEvent) PhotoID() string {
	return e.id
}
