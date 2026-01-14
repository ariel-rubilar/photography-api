package httperror

type Error struct {
	Err        error
	StatusCode int
}

func (e Error) Error() string {
	return e.Err.Error()
}

func Wrap(err error, statusCode int) *Error {
	return &Error{
		Err:        err,
		StatusCode: statusCode,
	}
}
