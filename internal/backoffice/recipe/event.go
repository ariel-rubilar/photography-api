package recipe

import "github.com/ariel-rubilar/photography-api/internal/shared/domain/event"

var (
	RecipeCreatedEventType event.Type = "backoffice.events.recipe.created"
)

type RecipeCreatedEventSettings struct {
	FilmSimulation       string
	DynamicRange         string
	Highlight            string
	Shadow               string
	Color                string
	NoiseReduction       string
	Sharpening           string
	Clarity              string
	GrainEffect          string
	ColorChromeEffect    string
	ColorChromeBlue      string
	WhiteBalance         string
	Iso                  string
	ExposureCompensation string
}

type RecipeCreatedEvent struct {
	event.BaseEvent
	id       string
	name     string
	settings RecipeCreatedEventSettings
	link     string
}

func newRecipeCreatedEvent(id, name, link string, settings RecipeCreatedEventSettings) RecipeCreatedEvent {
	return RecipeCreatedEvent{
		id:        id,
		name:      name,
		link:      link,
		settings:  settings,
		BaseEvent: event.NewBaseEvent(id),
	}
}

func (e RecipeCreatedEvent) Type() event.Type {
	return RecipeCreatedEventType
}

func (e RecipeCreatedEvent) RecipeID() string {
	return e.id
}

func (e RecipeCreatedEvent) Name() string {
	return e.name
}

func (e RecipeCreatedEvent) Settings() RecipeCreatedEventSettings {
	return e.settings
}

func (e RecipeCreatedEvent) Link() string {
	return e.link
}
