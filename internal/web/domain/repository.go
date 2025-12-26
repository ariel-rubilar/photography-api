package domain

import "context"

type Repository interface {
	Search(context.Context) ([]*Photo, error)
}
