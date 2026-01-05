package photoview

type Repository interface {
	Save(view *PhotoView) error
}
