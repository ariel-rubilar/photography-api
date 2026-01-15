package domainerror

type Error interface {
	error
	Code() string
}
