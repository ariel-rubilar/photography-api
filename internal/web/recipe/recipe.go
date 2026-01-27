package recipe

type RecipePrimitives struct {
	Name     string
	Settings RecipeSettingsPrimitives
	Link     string
	ID       string
}

type Recipe struct {
	id       string
	name     string
	settings RecipeSettings
	link     string
}

func new(id string, name string, link string, settings RecipeSettings) *Recipe {
	return &Recipe{
		id:       id,
		name:     name,
		settings: settings,
		link:     link,
	}
}

func Create(id string, name string, settings RecipeSettings, link string) (*Recipe, error) {
	r := new(id, name, link, settings)
	return r, nil
}

func FromPrimitives(rp RecipePrimitives) (*Recipe, error) {
	settings := RecipeSettingsFromPrimitives(rp.Settings)

	return new(rp.ID, rp.Name, rp.Link, settings), nil
}

func (r Recipe) ToPrimitives() RecipePrimitives {
	return RecipePrimitives{
		ID:       r.id,
		Name:     r.name,
		Settings: r.settings.ToPrimitives(),
		Link:     r.link,
	}
}

func (r Recipe) ID() string {
	return r.id
}
