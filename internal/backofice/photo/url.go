package photo

import "errors"

var (
	ErrInvalidPhotoUrl = errors.New("invalid photo url")
)

type PhotoUrl struct {
	value string
}

func NewPhotoUrl(value string) (PhotoUrl, error) {
	if value == "" {
		return PhotoUrl{}, ErrInvalidPhotoUrl
	}

	return PhotoUrl{value: value}, nil
}

func (pu PhotoUrl) Value() string {
	return pu.value
}
