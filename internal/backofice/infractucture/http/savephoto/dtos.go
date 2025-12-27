package savephoto

type photoDTO struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	URL      string `json:"url"`
	RecipeID string `json:"recipeId"`
}
