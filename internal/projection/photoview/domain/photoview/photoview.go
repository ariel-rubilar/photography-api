package photoview

type PhotoView struct {
	ID     string
	Title  string
	Url    string
	Recipe Recipe
}

func New(id, title, url string, recipe Recipe) *PhotoView {
	return &PhotoView{
		ID:     id,
		Title:  title,
		Url:    url,
		Recipe: recipe,
	}
}
