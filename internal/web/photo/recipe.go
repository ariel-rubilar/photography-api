package photo

type RecipePrimitives struct {
	Name     string
	Settings RecipeSettingsPrimitives
	Link     string
}

type Recipe struct {
	name     string
	settings RecipeSettings
	link     string
}

func newRecipe(name string, settings RecipeSettings, link string) Recipe {
	return Recipe{
		name:     name,
		settings: settings,
		link:     link,
	}
}

func RecipeFromPrimitives(rp RecipePrimitives) Recipe {
	return newRecipe(rp.Name, RecipeSettingsFromPrimitives(rp.Settings), rp.Link)
}

func (r Recipe) ToPrimitives() RecipePrimitives {
	return RecipePrimitives{
		Name:     r.name,
		Settings: r.settings.ToPrimitives(),
		Link:     r.link,
	}
}
