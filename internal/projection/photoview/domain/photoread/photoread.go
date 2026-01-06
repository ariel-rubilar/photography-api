package photoread

type PhotoRead struct {
	ID       string
	Title    string
	Url      string
	RecipeID string
}

func New(id, title, url, recipeID string) *PhotoRead {
	return &PhotoRead{
		ID:       id,
		Title:    title,
		Url:      url,
		RecipeID: recipeID,
	}
}
