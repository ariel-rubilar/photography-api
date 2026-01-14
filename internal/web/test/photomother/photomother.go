package photomother

import (
	"github.com/ariel-rubilar/photography-api/internal/web/photo"
	"github.com/ariel-rubilar/photography-api/internal/web/test/recipemother"
	"github.com/brianvoe/gofakeit/v7"
)

func DefaultPhoto() *photo.Photo {
	return photo.FromPrimitives(photo.PhotoPrimitives{
		ID:     gofakeit.UUID(),
		Title:  gofakeit.Sentence(3),
		URL:    gofakeit.URL(),
		Recipe: recipemother.DefaultRecipePrimitives(),
	})
}

type PhotoOption func(photo.PhotoPrimitives) photo.PhotoPrimitives

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

func WithRecipe(recipe photo.RecipePrimitives) PhotoOption {
	return func(p photo.PhotoPrimitives) photo.PhotoPrimitives {
		p.Recipe = recipe
		return p
	}
}

func NewPhoto(options ...PhotoOption) *photo.Photo {
	primitives := DefaultPhoto().ToPrimitives()
	for _, opt := range options {
		primitives = opt(primitives)
	}
	return photo.FromPrimitives(primitives)
}

func NewPhotoList(amount int) []*photo.Photo {
	photos := make([]*photo.Photo, amount)

	for i := range amount {
		photos[i] = NewPhoto()
	}

	return photos
}
