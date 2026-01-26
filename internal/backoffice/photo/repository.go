package photo

import "context"

type Repository interface {
	Save(context.Context, *Photo) error
	Exists(context.Context, string) (bool, error)
}
