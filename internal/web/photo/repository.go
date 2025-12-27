package photo

import "context"

type Repository interface {
	Search(context.Context) ([]*Photo, error)
}
