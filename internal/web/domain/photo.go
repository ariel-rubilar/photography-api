package domain

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

func NewPhoto(id, title, url string, recipe Recipe) *Photo {
	return &Photo{
		id:     id,
		title:  title,
		url:    url,
		recipe: recipe,
	}
}

func PhotoFromPrimitives(pr PhotoPrimitives) *Photo {
	return &Photo{
		id:     pr.ID,
		title:  pr.Title,
		url:    pr.URL,
		recipe: RecipeFromPrimitives(pr.Recipe),
	}
}

func (p *Photo) ToPrimitives() PhotoPrimitives {
	return PhotoPrimitives{
		ID:     p.id,
		Title:  p.title,
		URL:    p.url,
		Recipe: p.recipe.ToPrimitives(),
	}
}
