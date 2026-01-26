package recipedtomother

import (
	"github.com/ariel-rubilar/photography-api/internal/backoffice/usecases/recipequery"
	"github.com/brianvoe/gofakeit/v7"
)

type RecipeOption func(recipequery.RecipeDTO) recipequery.RecipeDTO

func DefaultRecipeDTOSettings() recipequery.RecipeDTOSettings {
	return recipequery.RecipeDTOSettings{
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

func DefaultRecipeDTO() recipequery.RecipeDTO {
	primitives := recipequery.RecipeDTO{
		ID:       gofakeit.UUID(),
		Name:     gofakeit.Sentence(2),
		Settings: DefaultRecipeDTOSettings(),
		Link:     gofakeit.URL(),
	}
	return primitives
}

func WithID(id string) RecipeOption {
	return func(p recipequery.RecipeDTO) recipequery.RecipeDTO {
		p.ID = id
		return p
	}
}

func WithName(name string) RecipeOption {
	return func(p recipequery.RecipeDTO) recipequery.RecipeDTO {
		p.Name = name
		return p
	}
}

func WithSettings(settings recipequery.RecipeDTOSettings) RecipeOption {
	return func(p recipequery.RecipeDTO) recipequery.RecipeDTO {
		p.Settings = settings
		return p
	}
}

func WithLink(link string) RecipeOption {
	return func(p recipequery.RecipeDTO) recipequery.RecipeDTO {
		p.Link = link
		return p
	}
}

func NewRecipeDTO(options ...RecipeOption) *recipequery.RecipeDTO {
	recipe := DefaultRecipeDTO()
	for _, opt := range options {
		recipe = opt(recipe)
	}
	return &recipe
}

func NewRecipeDTOList(amount int) []*recipequery.RecipeDTO {
	recipes := make([]*recipequery.RecipeDTO, amount)
	for i := range recipes {
		recipe := NewRecipeDTO()
		recipes[i] = recipe
	}
	return recipes
}
