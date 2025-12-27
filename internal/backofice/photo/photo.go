package photo

type PhotoPrimitives struct {
	ID       string
	Title    string
	URL      string
	RecipeID string
}

type Photo struct {
	id       string
	title    string
	url      PhotoUrl
	recipeID string
	events   []any
}

func new(id, title, recipeID string, url PhotoUrl) *Photo {
	return &Photo{
		id:       id,
		title:    title,
		url:      url,
		recipeID: recipeID,
	}
}

func Build(id, title, url, recipeID string) (*Photo, error) {
	urlVO, err := NewPhotoUrl(url)

	if err != nil {
		return nil, err
	}

	photo := new(
		id,
		title,
		recipeID,
		urlVO,
	)

	return photo, nil
}

func Create(id, title, url, recipeID string) (*Photo, error) {
	urlVO, err := NewPhotoUrl(url)

	if err != nil {
		return nil, err
	}

	photo := new(
		id,
		title,
		recipeID,
		urlVO,
	)

	photo.recordEvent(newPhotoCreatedEvent(id, recipeID))

	return photo, nil
}

func FromPrimitives(pr PhotoPrimitives) *Photo {
	return new(
		pr.ID,
		pr.Title,
		pr.RecipeID,
		PhotoUrl{
			value: pr.URL,
		},
	)
}

func (p *Photo) ToPrimitives() PhotoPrimitives {
	return PhotoPrimitives{
		ID:       p.id,
		Title:    p.title,
		URL:      p.url.Value(),
		RecipeID: p.recipeID,
	}
}

func (p *Photo) PullEvents() []any {
	events := p.events
	p.events = []any{}
	return events
}

func (p *Photo) recordEvent(event any) {
	p.events = append(p.events, event)
}
