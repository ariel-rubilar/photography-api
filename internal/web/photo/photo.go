package photo

type PhotoPrimitives struct {
	ID     string
	Title  string
	URL    string
	Recipe RecipePrimitives
}

type Photo struct {
	id     string
	title  string
	url    string
	recipe Recipe
}

func new(id, title, url string, recipe Recipe) *Photo {
	return &Photo{
		id:     id,
		title:  title,
		url:    url,
		recipe: recipe,
	}
}

func FromPrimitives(pr PhotoPrimitives) *Photo {
	return new(pr.ID, pr.Title, pr.URL, RecipeFromPrimitives(pr.Recipe))
}

func (p *Photo) ToPrimitives() PhotoPrimitives {
	return PhotoPrimitives{
		ID:     p.id,
		Title:  p.title,
		URL:    p.url,
		Recipe: p.recipe.ToPrimitives(),
	}
}
