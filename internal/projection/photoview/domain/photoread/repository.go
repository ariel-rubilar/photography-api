package photoread

import "context"

type Repository interface {
	Get(ctx context.Context, id string) (*PhotoRead, error)
}
