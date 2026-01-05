package photoread

type PhotoRead struct {
	ID       string
	Title    string
	Url      string
	RecipeID string
}

func new(id, title, recipeID string, url string) *PhotoRead {
	return &PhotoRead{
		ID:       id,
		Title:    title,
		Url:      url,
		RecipeID: recipeID,
	}
}

func Build(id, title, url, recipeID string) *PhotoRead {

	photo := new(
		id,
		title,
		recipeID,
		url,
	)

	return photo
}
