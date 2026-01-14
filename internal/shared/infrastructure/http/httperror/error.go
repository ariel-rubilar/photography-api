package httperror

type Error struct {
	Err        error
	StatusCode int
	Message    string
}

func (e Error) Error() string {
	return e.Err.Error()
}

func Wrap(err error, statusCode int, message string) *Error {
	return &Error{
		Err:        err,
		StatusCode: statusCode,
		Message:    message,
	}
}
