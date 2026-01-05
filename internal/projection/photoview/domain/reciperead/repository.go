package reciperead

type Repository interface {
	Get(id string) (*RecipeRead, error)
}
