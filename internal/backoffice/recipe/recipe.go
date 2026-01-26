package recipe

import "github.com/ariel-rubilar/photography-api/internal/shared/domain/event"

type RecipePrimitives struct {
	Name     string
	Settings RecipeSettingsPrimitives
	Link     string
	ID       string
}

type Recipe struct {
	id       string
	name     RecipeName
	settings RecipeSettings
	link     string
	events   []event.Event
}

func new(id string, name RecipeName, settings RecipeSettings, link string) *Recipe {
	return &Recipe{
		id:       id,
		name:     name,
		settings: settings,
		link:     link,
	}
}

func Create(id string, name string, settings RecipeSettings, link string) (*Recipe, error) {

	nameVO, err := NewRecipeName(name)
	if err != nil {
		return nil, err
	}

	r := new(id, nameVO, settings, link)

	r.recordEvent(newRecipeCreatedEvent(id, nameVO.value, link, RecipeCreatedEventSettings{
		FilmSimulation:       settings.filmSimulation,
		DynamicRange:         settings.dynamicRange,
		Highlight:            settings.highlight,
		Shadow:               settings.shadow,
		Color:                settings.color,
		NoiseReduction:       settings.noiseReduction,
		Sharpening:           settings.sharpening,
		Clarity:              settings.clarity,
		GrainEffect:          settings.grainEffect,
		ColorChromeEffect:    settings.colorChromeEffect,
		ColorChromeBlue:      settings.colorChromeBlue,
		WhiteBalance:         settings.whiteBalance,
		Iso:                  settings.iso,
		ExposureCompensation: settings.exposureCompensation,
	}))

	return r, nil
}

func FromPrimitives(rp RecipePrimitives) (*Recipe, error) {
	settings := RecipeSettingsFromPrimitives(rp.Settings)

	nameVO, err := NewRecipeName(rp.Name)
	if err != nil {
		return nil, err
	}

	return new(rp.ID, nameVO, settings, rp.Link), nil
}

func (r Recipe) ID() string {
	return r.id
}

func (r Recipe) Name() string {
	return r.name.Value()
}

func (r Recipe) ToPrimitives() RecipePrimitives {
	return RecipePrimitives{
		ID:       r.id,
		Name:     r.name.value,
		Settings: r.settings.ToPrimitives(),
		Link:     r.link,
	}
}

func (p *Recipe) recordEvent(event event.Event) {
	p.events = append(p.events, event)
}

func (r *Recipe) PullEvents() []event.Event {
	events := r.events
	r.events = []event.Event{}
	return events
}
