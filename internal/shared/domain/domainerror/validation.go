package domainerror

type Validation struct {
	Reason string
}

func (e Validation) Error() string {
	return e.Reason
}

func (e Validation) Code() string {
	return "VALIDATION_ERROR"
}
