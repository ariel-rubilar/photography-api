package photomother

import (
	"github.com/ariel-rubilar/photography-api/internal/backoffice/photo"
	"github.com/brianvoe/gofakeit/v7"
)

type PhotoOption func(photo.PhotoPrimitives) photo.PhotoPrimitives

func DefaultPhoto() *photo.Photo {
	primitives := photo.PhotoPrimitives{
		ID:       gofakeit.UUID(),
		Title:    gofakeit.Sentence(3),
		URL:      gofakeit.URL(),
		RecipeID: gofakeit.UUID(),
	}
	ph, _ := photo.FromPrimitives(primitives)
	return ph
}

func WithID(id string) PhotoOption {
	return func(p photo.PhotoPrimitives) photo.PhotoPrimitives {
		p.ID = id
		return p
	}
}

func WithTitle(title string) PhotoOption {
	return func(p photo.PhotoPrimitives) photo.PhotoPrimitives {
		p.Title = title
		return p
	}
}

func WithURL(url string) PhotoOption {
	return func(p photo.PhotoPrimitives) photo.PhotoPrimitives {
		p.URL = url
		return p
	}
}

func WithRecipeID(recipeID string) PhotoOption {
	return func(p photo.PhotoPrimitives) photo.PhotoPrimitives {
		p.RecipeID = recipeID
		return p
	}
}

func NewPhoto(options ...PhotoOption) *photo.Photo {
	primitives := DefaultPhoto().ToPrimitives()
	for _, opt := range options {
		primitives = opt(primitives)
	}
	ph, _ := photo.FromPrimitives(primitives)
	return ph
}

func NewPhotoList(amount int) []*photo.Photo {
	photos := make([]*photo.Photo, amount)
	for i := range photos {
		photos[i] = NewPhoto()
	}
	return photos
}
