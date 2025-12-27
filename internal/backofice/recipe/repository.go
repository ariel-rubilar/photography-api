package recipe

import "context"

type Repository interface {
	Search(context.Context) ([]*Recipe, error)
}
