package httperror

import "net/http"

type Error struct {
	Err        error
	StatusCode int
	Message    string
	Code       string
}

func (e Error) Error() string {
	return e.Err.Error()
}

type Options func(Error) Error

func Wrap(err error, statusCode int, options ...Options) Error {

	e := Error{
		Err:        err,
		StatusCode: http.StatusInternalServerError,
		Message:    "internal server error",
		Code:       "INTERNAL_ERROR",
	}

	for _, option := range options {
		e = option(e)
	}

	return e

}

func WithMessage(message string) Options {
	return func(e Error) Error {
		e.Message = message
		return e
	}
}

func WithCode(code string) Options {
	return func(e Error) Error {
		e.Code = code
		return e
	}
}

func WrapInternalServerError(err error, options ...Options) Error {
	return Wrap(err, http.StatusInternalServerError, options...)
}

func WrapBadRequestError(err error, options ...Options) Error {
	return Wrap(err, http.StatusInternalServerError, options...)
}
