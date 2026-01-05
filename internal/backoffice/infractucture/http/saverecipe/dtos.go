package saverecipe

import (
	"github.com/ariel-rubilar/photography-api/internal/backoffice/recipe"
)

type settingDTO struct {
	FilmSimulation       string `json:"filmSimulation"`
	DynamicRange         string `json:"dynamicRange"`
	Highlight            string `json:"highlight"`
	Shadow               string `json:"shadow"`
	Color                string `json:"color"`
	NoiseReduction       string `json:"noiseReduction"`
	Sharpening           string `json:"sharpening"`
	Clarity              string `json:"clarity"`
	GrainEffect          string `json:"grainEffect"`
	ColorChromeEffect    string `json:"colorChromeEffect"`
	ColorChromeBlue      string `json:"colorChromeBlue"`
	WhiteBalance         string `json:"whiteBalance"`
	Iso                  string `json:"iso"`
	ExposureCompensation string `json:"exposureCompensation"`
}

type recipeDTO struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Settings settingDTO `json:"settings"`
	Link     string     `json:"link"`
}

func (r *recipeDTO) toDomain() (*recipe.Recipe, error) {
	settings := recipe.NewRecipeSettings(
		r.Settings.FilmSimulation,
		r.Settings.DynamicRange,
		r.Settings.Highlight,
		r.Settings.Shadow,
		r.Settings.Color,
		r.Settings.NoiseReduction,
		r.Settings.Sharpening,
		r.Settings.Clarity,
		r.Settings.GrainEffect,
		r.Settings.ColorChromeEffect,
		r.Settings.ColorChromeBlue,
		r.Settings.WhiteBalance,
		r.Settings.Iso,
		r.Settings.ExposureCompensation,
	)

	return recipe.CreateRecipe(
		r.ID,
		r.Name,
		settings,
		r.Link,
	)
}
