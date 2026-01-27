package photodtomother

import (
	"github.com/ariel-rubilar/photography-api/internal/web/usecases/photoquery"
	"github.com/brianvoe/gofakeit/v7"
)

type PhotoOption func(photoquery.PhotoDTO) photoquery.PhotoDTO

func DefaultPhotoDTO() photoquery.PhotoDTO {
	primitives := photoquery.PhotoDTO{
		ID:     gofakeit.UUID(),
		Title:  gofakeit.Sentence(2),
		Recipe: DefaultPhotoDTORecipe(),
		URL:    gofakeit.URL(),
	}
	return primitives
}

func WithID(id string) PhotoOption {
	return func(p photoquery.PhotoDTO) photoquery.PhotoDTO {
		p.ID = id
		return p
	}
}

func WithTitle(value string) PhotoOption {
	return func(p photoquery.PhotoDTO) photoquery.PhotoDTO {
		p.Title = value
		return p
	}
}

func WithURL(value string) PhotoOption {
	return func(p photoquery.PhotoDTO) photoquery.PhotoDTO {
		p.URL = value
		return p
	}
}

func WithRecipe(value photoquery.PhotoDTORecipe) PhotoOption {
	return func(p photoquery.PhotoDTO) photoquery.PhotoDTO {
		p.Recipe = value
		return p
	}
}

func NewPhotoDTO(options ...PhotoOption) *photoquery.PhotoDTO {
	primitives := DefaultPhotoDTO()
	for _, opt := range options {
		primitives = opt(primitives)
	}
	return &primitives
}

func NewPhotoDTOList(amount int) []*photoquery.PhotoDTO {
	photos := make([]*photoquery.PhotoDTO, amount)
	for i := range photos {
		photos[i] = NewPhotoDTO()
	}
	return photos
}

type PhotoRecipeOption func(photoquery.PhotoDTORecipe) photoquery.PhotoDTORecipe

func DefaultPhotoDTORecipeSettings() photoquery.PhotoDTORecipeSettings {
	return photoquery.PhotoDTORecipeSettings{
		FilmSimulation:       gofakeit.Word(),
		DynamicRange:         gofakeit.Word(),
		Highlight:            gofakeit.Word(),
		Shadow:               gofakeit.Word(),
		Color:                gofakeit.Word(),
		NoiseReduction:       gofakeit.Word(),
		Sharpening:           gofakeit.Word(),
		Clarity:              gofakeit.Word(),
		GrainEffect:          gofakeit.Word(),
		ColorChromeEffect:    gofakeit.Word(),
		ColorChromeBlue:      gofakeit.Word(),
		WhiteBalance:         gofakeit.Word(),
		Iso:                  gofakeit.Word(),
		ExposureCompensation: gofakeit.Word(),
	}
}

func DefaultPhotoDTORecipe() photoquery.PhotoDTORecipe {
	primitives := photoquery.PhotoDTORecipe{
		ID:       gofakeit.UUID(),
		Name:     gofakeit.Sentence(2),
		Settings: DefaultPhotoDTORecipeSettings(),
		Link:     gofakeit.URL(),
	}
	return primitives
}

func WithPhotoDTORecipeID(id string) PhotoRecipeOption {
	return func(p photoquery.PhotoDTORecipe) photoquery.PhotoDTORecipe {
		p.ID = id
		return p
	}
}

func WithPhotoDTORecipeName(name string) PhotoRecipeOption {
	return func(p photoquery.PhotoDTORecipe) photoquery.PhotoDTORecipe {
		p.Name = name
		return p
	}
}

func WithPhotoDTORecipeSettings(settings photoquery.PhotoDTORecipeSettings) PhotoRecipeOption {
	return func(p photoquery.PhotoDTORecipe) photoquery.PhotoDTORecipe {
		p.Settings = settings
		return p
	}
}

func WithPhotoDTORecipeLink(link string) PhotoRecipeOption {
	return func(p photoquery.PhotoDTORecipe) photoquery.PhotoDTORecipe {
		p.Link = link
		return p
	}
}

func NewPhotoRecipeDTO(options ...PhotoRecipeOption) *photoquery.PhotoDTORecipe {
	recipe := DefaultPhotoDTORecipe()
	for _, opt := range options {
		recipe = opt(recipe)
	}
	return &recipe
}
