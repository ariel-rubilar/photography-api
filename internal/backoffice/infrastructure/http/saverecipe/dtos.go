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
	settings := recipe.RecipeSettingsFromPrimitives(
		recipe.RecipeSettingsPrimitives{
			FilmSimulation:       r.Settings.FilmSimulation,
			DynamicRange:         r.Settings.DynamicRange,
			Highlight:            r.Settings.Highlight,
			Shadow:               r.Settings.Shadow,
			Color:                r.Settings.Color,
			NoiseReduction:       r.Settings.NoiseReduction,
			Sharpening:           r.Settings.Sharpening,
			Clarity:              r.Settings.Clarity,
			GrainEffect:          r.Settings.GrainEffect,
			ColorChromeEffect:    r.Settings.ColorChromeEffect,
			ColorChromeBlue:      r.Settings.ColorChromeBlue,
			WhiteBalance:         r.Settings.WhiteBalance,
			Iso:                  r.Settings.Iso,
			ExposureCompensation: r.Settings.ExposureCompensation,
		},
	)

	return recipe.Create(
		r.ID,
		r.Name,
		settings,
		r.Link,
	)
}
