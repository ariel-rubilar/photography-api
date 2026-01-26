package recipe

import "context"

type Repository interface {
	Search(context.Context, Criteria) ([]*Recipe, error)
	Save(context.Context, *Recipe) error
	Exists(ctx context.Context, id string) (bool, error)
}
