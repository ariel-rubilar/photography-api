package error

type Conflict struct {
	Reason string
}

func (e Conflict) Error() string {
	return e.Reason
}

func (e Conflict) Code() string {
	return "CONFLICT"
}
