package recipe

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

	return new(id, nameVO, settings, link), nil
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
