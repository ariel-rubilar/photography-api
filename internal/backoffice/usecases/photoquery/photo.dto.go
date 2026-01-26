package photoquery

type PhotoDTO struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	URL      string `json:"url"`
	RecipeID string `json:"recipe_id"`
}
