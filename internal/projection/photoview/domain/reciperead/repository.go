package reciperead

import "context"

type Repository interface {
	Get(ctx context.Context, id string) (*RecipeRead, error)
}
