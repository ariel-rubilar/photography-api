package photo

type RecipePrimitives struct {
	ID       string
	Name     string
	Settings RecipeSettingsPrimitives
	Link     string
}

type Recipe struct {
	id       string
	name     string
	settings RecipeSettings
	link     string
}

func newRecipe(id string, name string, settings RecipeSettings, link string) Recipe {
	return Recipe{
		id:       id,
		name:     name,
		settings: settings,
		link:     link,
	}
}

func RecipeFromPrimitives(rp RecipePrimitives) Recipe {
	return newRecipe(rp.ID, rp.Name, RecipeSettingsFromPrimitives(rp.Settings), rp.Link)
}

func (r Recipe) ToPrimitives() RecipePrimitives {
	return RecipePrimitives{
		ID:       r.id,
		Name:     r.name,
		Settings: r.settings.ToPrimitives(),
		Link:     r.link,
	}
}
