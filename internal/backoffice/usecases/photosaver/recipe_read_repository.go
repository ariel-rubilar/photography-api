package photosaver

import "context"

type RecipeReadRepository interface {
	Exists(ctx context.Context, id string) (bool, error)
}
