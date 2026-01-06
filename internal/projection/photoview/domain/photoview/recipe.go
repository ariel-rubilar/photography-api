package photoview

type Recipe struct {
	ID       string
	Name     string
	Settings RecipeSettings
	Link     string
}

func NewRecipe(id, name string, settings RecipeSettings, link string) Recipe {
	return Recipe{
		ID:       id,
		Name:     name,
		Settings: settings,
		Link:     link,
	}
}
