package recipe

import "context"

type Repository interface {
	Save(context.Context, *Recipe) error
	Exists(ctx context.Context, id string) (bool, error)
}
