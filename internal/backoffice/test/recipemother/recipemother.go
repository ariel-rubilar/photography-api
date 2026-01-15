package recipemother

import (
	"github.com/ariel-rubilar/photography-api/internal/backoffice/recipe"
	"github.com/brianvoe/gofakeit/v7"
)

type RecipeOption func(recipe.RecipePrimitives) recipe.RecipePrimitives

func DefaultRecipeSettingsPrimitives() recipe.RecipeSettingsPrimitives {
	return recipe.RecipeSettingsPrimitives{
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

func DefaultRecipe() *recipe.Recipe {
	primitives := recipe.RecipePrimitives{
		ID:       gofakeit.UUID(),
		Name:     gofakeit.Sentence(2),
		Settings: DefaultRecipeSettingsPrimitives(),
		Link:     gofakeit.URL(),
	}
	r, _ := recipe.FromPrimitives(primitives)
	return r
}

func WithID(id string) RecipeOption {
	return func(p recipe.RecipePrimitives) recipe.RecipePrimitives {
		p.ID = id
		return p
	}
}

func WithName(name string) RecipeOption {
	return func(p recipe.RecipePrimitives) recipe.RecipePrimitives {
		p.Name = name
		return p
	}
}

func WithSettings(settings recipe.RecipeSettingsPrimitives) RecipeOption {
	return func(p recipe.RecipePrimitives) recipe.RecipePrimitives {
		p.Settings = settings
		return p
	}
}

func WithLink(link string) RecipeOption {
	return func(p recipe.RecipePrimitives) recipe.RecipePrimitives {
		p.Link = link
		return p
	}
}

func NewRecipe(options ...RecipeOption) *recipe.Recipe {
	primitives := DefaultRecipe().ToPrimitives()
	for _, opt := range options {
		primitives = opt(primitives)
	}
	r, _ := recipe.FromPrimitives(primitives)
	return r
}

func NewRecipeList(amount int) []*recipe.Recipe {
	recipes := make([]*recipe.Recipe, amount)
	for i := range recipes {
		recipes[i] = NewRecipe()
	}
	return recipes
}
