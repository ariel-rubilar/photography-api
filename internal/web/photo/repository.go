package photo

import "context"

type Repository interface {
	Search(context.Context) ([]*Photo, error)
	Save(context.Context, *Photo) error
	Exists(ctx context.Context, id string) (bool, error)
}
