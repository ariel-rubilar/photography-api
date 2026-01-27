package photo

import "github.com/ariel-rubilar/photography-api/internal/shared/domain/event"

var (
	PhotoCreatedEventType event.Type = "backoffice.events.photo.created"
)

type PhotoCreatedEvent struct {
	event.BaseEvent
	id       string
	recipeID string
	title    string
	url      string
}

func newPhotoCreatedEvent(id, recipeID, title, url string) PhotoCreatedEvent {
	return PhotoCreatedEvent{
		id:        id,
		recipeID:  recipeID,
		BaseEvent: event.NewBaseEvent(),
		title:     title,
		url:       url,
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

func (e PhotoCreatedEvent) Title() string {
	return e.title
}

func (e PhotoCreatedEvent) URL() string {
	return e.url
}
