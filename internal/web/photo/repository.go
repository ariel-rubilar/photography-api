package photo

import "context"

type Repository interface {
	Save(context.Context, *Photo) error
	Exists(ctx context.Context, id string) (bool, error)
}
