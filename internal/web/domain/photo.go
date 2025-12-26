package domain

type Photo struct {
	ID    string
	Title string
	URL   string
}

func NewPhoto(id, title, url string) *Photo {
	return &Photo{
		ID:    id,
		Title: title,
		URL:   url,
	}
}
