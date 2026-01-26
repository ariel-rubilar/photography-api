package photoquery

import "context"

type Repository interface {
	Search(context.Context, Criteria) ([]*PhotoDTO, error)
}
