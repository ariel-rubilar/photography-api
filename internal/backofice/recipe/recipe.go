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

func NewRecipe(id string, name string, settings RecipeSettings, link string) *Recipe {
	return &Recipe{
		id:       id,
		name:     name,
		settings: settings,
		link:     link,
	}
}

func CreateRecipe(id string, name string, settings RecipeSettings, link string) (*Recipe, error) {
	return &Recipe{
		id:       id,
		name:     name,
		settings: settings,
		link:     link,
	}, nil
}

func RecipeFromPrimitives(rp RecipePrimitives) Recipe {
	return Recipe{
		name:     rp.Name,
		settings: RecipeSettingsFromPrimitives(rp.Settings),
		link:     rp.Link,
	}
}

func (r Recipe) ToPrimitives() RecipePrimitives {
	return RecipePrimitives{
		ID:       r.id,
		Name:     r.name,
		Settings: r.settings.ToPrimitives(),
		Link:     r.link,
	}
}
