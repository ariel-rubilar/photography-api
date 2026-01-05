package photoread

type Repository interface {
	Get(id string) (*PhotoRead, error)
}
