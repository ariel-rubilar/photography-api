package photo

import "context"

type Repository interface {
	Save(context.Context, *Photo) error
	Search(context.Context, Criteria) ([]*Photo, error)
}
