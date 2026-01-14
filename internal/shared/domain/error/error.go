package error

type Error interface {
	error
	Code() string
}
