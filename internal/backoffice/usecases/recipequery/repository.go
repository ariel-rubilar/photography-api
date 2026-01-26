package recipequery

import "context"

type Repository interface {
	Search(context.Context, Criteria) ([]*RecipeDTO, error)
}
