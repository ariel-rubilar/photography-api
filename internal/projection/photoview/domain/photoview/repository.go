package photoview

import "context"

type Repository interface {
	Save(ctx context.Context, view *PhotoView) error
}
