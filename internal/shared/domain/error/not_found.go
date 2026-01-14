package error

type NotFound struct {
	Entity string
}

func (e NotFound) Error() string {
	return e.Entity + " not found"
}

func (e NotFound) Code() string {
	return "NOT_FOUND"
}
